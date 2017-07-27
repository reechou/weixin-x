package controller

import (
	"sync"
	"time"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
)

type WeixinMonitorManager struct {
	sync.Mutex

	weixinMonitorMap map[int64]*WeixinMonitor

	stop chan struct{}
}

func NewWeixinMonitorManager() *WeixinMonitorManager {
	wmm := &WeixinMonitorManager{
		weixinMonitorMap: make(map[int64]*WeixinMonitor),
		stop:             make(chan struct{}),
	}
	wmm.getLiebianTypeList()
	
	go wmm.loopGetLiebianTypeList()

	return wmm
}

func (self *WeixinMonitorManager) stop() {
	close(self.stop)
	for _, v := range self.weixinMonitorMap {
		v.Stop()
	}
}

func (self *WeixinMonitorManager) loopGetLiebianTypeList() {
	holmes.Info("loop get liebian type list started.")
	for {
		select {
		case <-time.After(10 * time.Minute):
			self.getLiebianTypeList()
		case <-self.stop:
			return
		}
	}
}

func (self *WeixinMonitorManager) getLiebianTypeList() {
	self.Lock()
	defer self.Unlock()

	liebianTypeList, err := models.GetLiebianTypeList()
	if err != nil {
		holmes.Error("get liebian type list error: %v", err)
		return
	}

	for _, v := range liebianTypeList {
		_, ok := self.weixinMonitorMap[v.LiebianType]
		if !ok {
			weixinMonitor := NewWeixinMonitor(v.LiebianType)
			self.weixinMonitorMap[v.LiebianType] = weixinMonitor
		}
	}
}
