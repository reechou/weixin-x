package controller

import (
	"strings"
	"time"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
	"github.com/reechou/weixin-x/models"
	"github.com/robfig/cron"
	//"github.com/jinzhu/now"
)

type TimerInfo struct {
	BeforeStart int64
	BeforeEnd   int64
	Desc        string
}

//var TimerInfoMap = map[int64]*TimerInfo{
//	1:  &TimerInfo{BeforeStart: 3600, BeforeEnd: 1800, Desc: "往前一小时-往前半小时"},
//	2:  &TimerInfo{BeforeStart: 5400, BeforeEnd: 3600, Desc: "往前一个半小时-往前一小时"},
//	3:  &TimerInfo{BeforeStart: 7200, BeforeEnd: 5400, Desc: "往前二小时-往前一个半小时"},
//	4:  &TimerInfo{BeforeStart: 9000, BeforeEnd: 7200, Desc: "往前二个半小时-往前二小时"},
//	5:  &TimerInfo{BeforeStart: 10800, BeforeEnd: 9000, Desc: "往前三小时-往前二个半小时"},
//	6:  &TimerInfo{BeforeStart: 12600, BeforeEnd: 10800, Desc: "往前三个半小时-往前三小时"},
//	7:  &TimerInfo{BeforeStart: 14400, BeforeEnd: 12600, Desc: "往前四小时-往前三个半小时"},
//	8:  &TimerInfo{BeforeStart: 16200, BeforeEnd: 14400, Desc: "往前四个半小时-往前四小时"},
//	9:  &TimerInfo{BeforeStart: 18000, BeforeEnd: 16200, Desc: "往前五小时-往前四个半小时"},
//	10: &TimerInfo{BeforeStart: 19800, BeforeEnd: 18000, Desc: "往前五个半小时-往前五小时"},
//	11: &TimerInfo{BeforeStart: 21600, BeforeEnd: 19800, Desc: "往前六小时-往前五个半小时"},
//	12: &TimerInfo{BeforeStart: 23400, BeforeEnd: 21600, Desc: "往前六个半小时-往前六小时"},
//	13: &TimerInfo{BeforeStart: 25200, BeforeEnd: 23400, Desc: "往前七小时-往前六个半小时"},
//	14: &TimerInfo{BeforeStart: 27000, BeforeEnd: 25200, Desc: "往前七个半小时-往前七小时"},
//	15: &TimerInfo{BeforeStart: 28800, BeforeEnd: 27000, Desc: "往前八小时-往前七个半小时"},
//	16: &TimerInfo{BeforeStart: 30600, BeforeEnd: 28800, Desc: "往前八个半小时-往前八小时"},
//	17: &TimerInfo{BeforeStart: 32400, BeforeEnd: 30600, Desc: "往前九小时-往前八个半小时"},
//	18: &TimerInfo{BeforeStart: 34200, BeforeEnd: 32400, Desc: "往前九个半小时-往前九小时"},
//	19: &TimerInfo{BeforeStart: 36000, BeforeEnd: 34200, Desc: "往前十小时-往前九个半小时"},
//	20: &TimerInfo{BeforeStart: 37800, BeforeEnd: 36000, Desc: "往前十个半小时-往前十小时"},
//	21: &TimerInfo{BeforeStart: 39600, BeforeEnd: 37800, Desc: "往前十一小时-往前十个半小时"},
//	22: &TimerInfo{BeforeStart: 41400, BeforeEnd: 39600, Desc: "往前十一个半小时-往前十一小时"},
//	23: &TimerInfo{BeforeStart: 43200, BeforeEnd: 41400, Desc: "往前十二小时-往前十一个半小时"},
//	24: &TimerInfo{BeforeStart: 45000, BeforeEnd: 43200, Desc: "往前十二个半小时-往前十二小时"},
//}

var TimerInfoMap = map[int64]*TimerInfo{
	1:  &TimerInfo{BeforeStart: 5400, BeforeEnd: 1800, Desc: "往前一半小时-往前半小时"},
	2:  &TimerInfo{BeforeStart: 9000, BeforeEnd: 5400, Desc: "往前二个半个半小时-往前一半小时"},
	3:  &TimerInfo{BeforeStart: 12600, BeforeEnd: 9000, Desc: "往前三个半小时-往前二个半小时"},
	4:  &TimerInfo{BeforeStart: 16200, BeforeEnd: 12600, Desc: "往前四个半小时-往前三个半小时"},
	5:  &TimerInfo{BeforeStart: 19800, BeforeEnd: 16200, Desc: "往前五个半小时-往前四个半小时"},
	6:  &TimerInfo{BeforeStart: 23400, BeforeEnd: 19800, Desc: "往前六个半小时-往前五个半小时"},
	7:  &TimerInfo{BeforeStart: 27000, BeforeEnd: 23400, Desc: "往前七个半小时-往前六个半小时"},
	8:  &TimerInfo{BeforeStart: 30600, BeforeEnd: 27000, Desc: "往前八个半小时-往前七个半小时"},
	9:  &TimerInfo{BeforeStart: 34200, BeforeEnd: 30600, Desc: "往前九个半小时-往前八个半小时"},
	10: &TimerInfo{BeforeStart: 37800, BeforeEnd: 34200, Desc: "往前十个半小时-往前九个半小时"},
	11: &TimerInfo{BeforeStart: 41400, BeforeEnd: 37800, Desc: "往前十一个半小时-往前十个半小时"},
	12: &TimerInfo{BeforeStart: 45000, BeforeEnd: 41400, Desc: "往前十二个半小时-往前十一个半小时"},
}

type TimerTaskWorker struct {
	cfg *config.Config

	stop chan struct{}
}

func NewTimerTaskWorker(cfg *config.Config) *TimerTaskWorker {
	ttw := &TimerTaskWorker{
		cfg:  cfg,
		stop: make(chan struct{}),
	}

	go ttw.start()

	return ttw
}

func (self *TimerTaskWorker) Stop() {
	close(self.stop)
}

func (self *TimerTaskWorker) start() {
	holmes.Debug("timer task worker cron start.")

	c := cron.New()
	c.AddFunc(self.cfg.TimerTaskCron, self.runTimerTask)
	c.Start()

	select {
	case <-self.stop:
		c.Stop()
		return
	}
}

func (self *TimerTaskWorker) runTimerTask() {
	holmes.Debug("[timer task] run start.")
	start := time.Now()
	defer func() {
		holmes.Debug("[timer task] run end, user time: %v.", time.Now().Sub(start))
	}()

	minute := start.Minute()
	//minute := 10

	timerTaskList, err := models.GetTimerTaskList()
	if err != nil {
		holmes.Error("get timer task list error: %v", err)
		return
	}
	//nowTime := now.BeginningOfHour().Unix()
	nowTime := start.Unix()
	for _, v := range timerTaskList {
		modId := int(v.WeixinId) % 60
		if modId != minute {
			continue
		}

		timerInfo, ok := TimerInfoMap[v.TimeId]
		if !ok {
			holmes.Error("timertask[%v] cannot found timerinfo", v)
			continue
		}
		startTime := nowTime - timerInfo.BeforeStart
		endTime := nowTime - timerInfo.BeforeEnd
		var friendsList []string
		if v.TagId != 0 {
			friends, err := models.GetWxTagFriendList(v.WeixinId, v.TagId, startTime, endTime)
			if err != nil {
				holmes.Error("get tag friend error: %v", err)
				continue
			}
			for _, v2 := range friends {
				friendsList = append(friendsList, v2.WeixinContact.UserName)
			}
		} else {
			friends, err := models.GetWeixinContactListFromTime(v.WeixinId, startTime, endTime)
			if err != nil {
				holmes.Error("get weixin contact from time error: %v", err)
				continue
			}
			for _, v2 := range friends {
				friendsList = append(friendsList, v2.UserName)
			}
		}
		//holmes.Debug("friends: %d %d %d %v", nowTime, startTime, endTime, friendsList)
		if len(friendsList) != 0 {
			task := &models.WeixinTaskList{
				WeixinId:     v.WeixinId,
				WeixinTaskId: v.TaskId,
				Friends:      strings.Join(friendsList, ","),
			}
			err = models.CreateWeixinTaskList(task)
			if err != nil {
				holmes.Error("create weixin task list error: %v", err)
			}
		}
	}
}
