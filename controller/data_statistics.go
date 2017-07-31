package controller
//
//import (
//	"strings"
//	"time"
//	"sync"
//
//	"github.com/reechou/holmes"
//	"github.com/reechou/weixin-x/config"
//	"github.com/reechou/weixin-x/models"
//	"github.com/robfig/cron"
//)
//
//type StatisticsDataInfo struct {
//	TypeId int64
//	Data   int64
//}
//
//type DataStatisticsWorker struct {
//	sync.Mutex
//
//	cfg *config.Config
//
//	stop chan struct{}
//}
//
//func (self *DataStatisticsWorker) start() {
//	holmes.Debug("data statistics worker cron start.")
//
//	c := cron.New()
//	c.AddFunc(self.cfg.TimerTaskCron, self.runSave)
//	c.Start()
//
//	select {
//	case <-self.stop:
//		c.Stop()
//		return
//	}
//}
//
//func (self *DataStatisticsWorker) runSave() {
//
//}
