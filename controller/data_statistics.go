package controller

import (
	"sync"

	"github.com/jinzhu/now"
	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
	"github.com/reechou/weixin-x/models"
	"github.com/robfig/cron"
)

type StatisticsDataInfo struct {
	TypeId int64
	Data   int64
}

type DataStatisticsWorker struct {
	sync.Mutex

	cfg      *config.Config
	dataChan chan *StatisticsDataInfo
	dataMap  map[int64]int64
	nowHour  int64

	stop chan struct{}
}

func NewDataStatisticsWorker(cfg *config.Config) *DataStatisticsWorker {
	dsw := &DataStatisticsWorker{
		cfg:      cfg,
		dataChan: make(chan *StatisticsDataInfo, 1024),
		dataMap:  make(map[int64]int64),
		stop:     make(chan struct{}),
	}
	dsw.init()
	
	go dsw.runCollect()
	go dsw.start()
	
	return dsw
}

func (self *DataStatisticsWorker) init() {
	self.Lock()
	defer self.Unlock()

	self.nowHour = now.BeginningOfHour().Unix()
	bindCount, err := models.GetBindCardCountFromTime(self.nowHour, 0)
	if err != nil {
		holmes.Error("get bind cart count error: %v", err)
	}
	self.dataMap[int64(models.S_DATA_SCREENSHOT)] = bindCount
	contactCount, err := models.GetWeixinContactCountFromTime(self.nowHour, 0)
	if err != nil {
		holmes.Error("get contact count from time error: %v", err)
	}
	self.dataMap[int64(models.S_DATA_ADD_CONTACT)] = contactCount
	
	holmes.Debug("now hour: %d data map: %v", self.nowHour, self.dataMap)
}

func (self *DataStatisticsWorker) Collect(data *StatisticsDataInfo) {
	select {
	case self.dataChan <- data:
	case <-self.stop:
		return
	}
}

func (self *DataStatisticsWorker) runCollect() {
	for {
		select {
		case data := <-self.dataChan:
			self.saveData(data)
		case <-self.stop:
			return
		}
	}
}

func (self *DataStatisticsWorker) saveData(data *StatisticsDataInfo) {
	self.Lock()
	defer self.Unlock()

	v, ok := self.dataMap[data.TypeId]
	if ok {
		self.dataMap[data.TypeId] = v + data.Data
	}
}

func (self *DataStatisticsWorker) start() {
	holmes.Debug("data statistics worker cron start.")

	c := cron.New()
	c.AddFunc(self.cfg.DataCollectCron, self.runSave)
	c.Start()

	select {
	case <-self.stop:
		c.Stop()
		return
	}
}

func (self *DataStatisticsWorker) runSave() {
	nowHour := now.BeginningOfHour().Unix()
	dataMap := make(map[int64]int64)
	self.Lock()
	for k, v := range self.dataMap {
		dataMap[k] = v
	}
	self.Unlock()
	for k, v := range dataMap {
		sd := &models.StatisticalData{
			TypeId:     k,
			Data:       v,
			TimeSeries: self.nowHour,
		}
		has, err := models.GetStatisticalData(&models.StatisticalData{TypeId: k, TimeSeries: self.nowHour})
		if err != nil {
			holmes.Error("get statistical data error: %v", err)
			continue
		}
		holmes.Debug("run save: %v %v", has, sd)
		if has {
			_, err := models.UpdateStatisticalData(sd)
			if err != nil {
				holmes.Error("update statistical data error: %v", err)
				continue
			}
		} else {
			err = models.CreateStatisticalData(sd)
			if err != nil {
				holmes.Error("create statistical data error: %v", err)
				continue
			}
		}
		if nowHour != self.nowHour {
			self.Lock()
			self.dataMap[k] = 0
			self.Unlock()
		}
	}
	if nowHour != self.nowHour {
		self.nowHour = nowHour
	}
}
