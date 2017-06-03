package controller

import (
	"sync"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/proto"
)

const (
	RUN_WORKER_LEN = 1024
)

type ContactWorker struct {
	wg sync.WaitGroup

	contactChan chan *proto.SyncContacts

	stop chan struct{}
}

func NewContactWorker() *ContactWorker {
	cw := &ContactWorker{
		contactChan: make(chan *proto.SyncContacts, 1024),
		stop:        make(chan struct{}),
	}

	for i := 0; i < RUN_WORKER_LEN; i++ {
		cw.wg.Add(1)
		go cw.runWorker()
	}
	
	holmes.Debug("contact workers started.")

	return cw
}

func (self *ContactWorker) Stop() {
	close(self.stop)
	self.wg.Wait()
}

func (self *ContactWorker) Sync(c *proto.SyncContacts) {
	select {
	case self.contactChan <- c:
	case <-self.stop:
		return
	}
}

func (self *ContactWorker) runWorker() {
	for {
		select {
		case c := <-self.contactChan:
			self.runSyncContacts(c)
		case <-self.stop:
			self.wg.Done()
			return
		}
	}
}

func (self *ContactWorker) runSyncContacts(c *proto.SyncContacts) {
	weixin := &models.Weixin{
		WxId: c.Myself,
	}
	has, err := models.GetWeixinFromWxid(weixin)
	if err != nil {
		holmes.Error("get weixin error: %v", err)
		return
	}
	if !has {
		holmes.Error("Has no this wxid[%s]", c.Myself)
		return
	}
	
	holmes.Debug("wechat[%s] nick[%s] run sync contacts starting.", weixin.Wechat, weixin.NickName)

	for _, v := range c.ContactData {
		if v.UserName == "" {
			continue
		}
		wc := &models.WeixinContact{
			UserName: v.UserName,
		}
		has, err := models.GetWeixinContact(wc)
		if err != nil {
			holmes.Error("get weixin contact error: %v", err)
			continue
		}
		if has {
			continue
		}
		wc = &models.WeixinContact{
			WeixinId:    weixin.ID,
			UserName:    v.UserName,
			AliasName:   v.AliasName,
			NickName:    v.NickName,
			PhoneNumber: v.PhoneNumber,
			Country:     v.Country,
			Province:    v.Province,
			City:        v.City,
			Sex:         v.Sex,
			Remark:      v.Remark,
		}
		err = models.CreateWeixinContact(wc)
		if err != nil {
			holmes.Error("create weixin contact error: %v", err)
		}
	}
	
	holmes.Debug("wechat[%s] nick[%s] run sync contacts end.", weixin.Wechat, weixin.NickName)
}
