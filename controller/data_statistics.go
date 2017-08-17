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
	TypeId      int64
	Data        int64
	LiebianType int64
	WeixinId    int64
	WxId        string
}

type DataStatisticsWorker struct {
	dataLock sync.Mutex

	cfg      *config.Config
	dataChan chan *StatisticsDataInfo
	dataMap  map[int64]map[int64]int64
	nowHour  int64

	stop chan struct{}
}

func NewDataStatisticsWorker(cfg *config.Config) *DataStatisticsWorker {
	dsw := &DataStatisticsWorker{
		cfg:      cfg,
		dataChan: make(chan *StatisticsDataInfo, 1024),
		dataMap:  make(map[int64]map[int64]int64),
		stop:     make(chan struct{}),
	}
	dsw.init()

	go dsw.runCollect()
	go dsw.start()

	return dsw
}

func (self *DataStatisticsWorker) init() {
	self.dataLock.Lock()
	defer self.dataLock.Unlock()
	
	self.nowHour = now.BeginningOfHour().Unix()
	
	self.dataMap[int64(models.S_DATA_SCREENSHOT)] = make(map[int64]int64)
	self.dataMap[int64(models.S_DATA_ADD_CONTACT)] = make(map[int64]int64)
	self.dataMap[int64(models.S_DATA_LIEBIAN_PV)] = make(map[int64]int64)
	self.dataMap[int64(models.S_DATA_LIEBIAN_UV)] = make(map[int64]int64)

	//bindCount, err := models.GetBindCardCountFromTime(self.nowHour, 0)
	//if err != nil {
	//	holmes.Error("get bind cart count error: %v", err)
	//}
	//self.dataMap[int64(models.S_DATA_SCREENSHOT)] = bindCount
	//contactCount, err := models.GetWeixinContactCountFromTime(self.nowHour, 0)
	//if err != nil {
	//	holmes.Error("get contact count from time error: %v", err)
	//}
	//self.dataMap[int64(models.S_DATA_ADD_CONTACT)] = contactCount
	//
	//holmes.Debug("now hour: %d data map: %v", self.nowHour, self.dataMap)
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
	var liebianType int64
	if data.LiebianType != 0 {
		liebianType = data.LiebianType
	} else {
		if data.WeixinId != 0 {
			lp := &models.LiebianPool{
				WeixinId: data.WeixinId,
			}
			has, err := models.GetLiebianPool(lp)
			if err != nil {
				holmes.Error("get liebian pool error: %v", err)
				return
			}
			if !has {
				holmes.Error("cannot found this weixin[%d] in liebian pool", data.WeixinId)
				return
			}
			liebianType = lp.LiebianType
		} else {
			if data.WxId == "" {
				holmes.Error("weixin id or wxid cannot be nil")
				return
			}
			liebianInfo, err := models.GetLiebianPoolFromWxId(data.WxId)
			if err != nil {
				holmes.Error("get liebian pool from wxid error: %v", err)
				return
			}
			if liebianInfo == nil {
				holmes.Error("get liebian pool from wxid[%s] is nil", data.WxId)
				return
			}
			liebianType = liebianInfo.LiebianPool.LiebianType
		}
	}

	self.dataLock.Lock()
	defer self.dataLock.Unlock()

	v, ok := self.dataMap[data.TypeId]
	if ok {
		v2, ok2 := v[liebianType]
		if ok2 {
			v[liebianType] = v2 + data.Data
		} else {
			v[liebianType] = data.Data
		}
		//self.dataMap[data.TypeId] = v + data.Data
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
	//holmes.Debug("start run save: %v", self.dataMap)
	nowHour := now.BeginningOfHour().Unix()
	dataMap := make(map[int64]map[int64]int64)
	self.dataLock.Lock()
	for k, v := range self.dataMap {
		dataMap[k] = make(map[int64]int64)
		for k2, v2 := range v {
			dataMap[k][k2] = v2
		}
	}
	self.dataLock.Unlock()
	//holmes.Debug("start data map: %v", dataMap)
	for k, v := range dataMap {
		for k2, v2 := range v {
			sd := &models.StatisticalData{
				TypeId:      k,
				Data:        v2,
				TimeSeries:  self.nowHour,
				LiebianType: k2,
			}
			has, err := models.GetStatisticalData(&models.StatisticalData{TypeId: k, TimeSeries: self.nowHour, LiebianType: k2})
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
				self.dataLock.Lock()
				self.dataMap[k][k2] = 0
				self.dataLock.Unlock()
			}
		}
	}
	if nowHour != self.nowHour {
		self.nowHour = nowHour
	}
}
