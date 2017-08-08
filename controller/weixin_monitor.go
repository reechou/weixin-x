package controller

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
)

const (
	CHECK_HEALTH_TIME = 300
)

const (
	MONITOR_ABNORMAL_WARN_MSG = "[微信监控] 裂变Type[%d] 机器 %v 出现异常，5分钟内未上报心跳，已自动下线。"
)

type WeixinMonitor struct {
	LiebianType int64

	stop chan struct{}
	done chan struct{}
}

func NewWeixinMonitor(liebianType int64) *WeixinMonitor {
	wm := &WeixinMonitor{
		LiebianType: liebianType,
		stop:        make(chan struct{}),
		done:        make(chan struct{}),
	}

	go wm.run()

	return wm
}

func (self *WeixinMonitor) Stop() {
	close(self.stop)
	<-self.done
}

func (self *WeixinMonitor) run() {
	holmes.Info("liebian[%d] monitor check has started.", self.LiebianType)
	for {
		select {
		case <-time.After(time.Minute):
			self.check()
		case <-self.stop:
			close(self.done)
			return
		}
	}
}

func (self *WeixinMonitor) check() {
	weixinList, err := models.GetLiebianPoolWeixinList(self.LiebianType)
	if err != nil {
		holmes.Error("get liebian pool weixin list error: %v", err)
		return
	}
	healthNode := 0
	now := time.Now().Unix()
	var abnormalIds []int64
	for _, v := range weixinList {
		if now-v.Weixin.LastHeartbeat < CHECK_HEALTH_TIME {
			healthNode++
			continue
		}
		holmes.Error("weixin[%v] check health error.", v)
		// down this weixin
		err = models.DelLiebianPoolList([]int64{v.LiebianPool.ID})
		if err != nil {
			holmes.Error("del liebian pool[%d] error: %v", v.LiebianPool.ID, err)
		}
		abnormalIds = append(abnormalIds, v.Weixin.ID)
	}
	// send warn msg
	if len(abnormalIds) > 0 {
		abnormalMsg := fmt.Sprintf(MONITOR_ABNORMAL_WARN_MSG, self.LiebianType, abnormalIds)
		holmes.Error("abnormal msg: %s", abnormalMsg)
		errorMsg := models.LiebianErrorMsg{
			LiebianType: self.LiebianType,
			Msg:         abnormalMsg,
		}
		models.CreateLiebianErrorMsg(errorMsg)
	}
}
