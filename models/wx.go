package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

const (
	WX_TYPE_WECHAT = iota
	WX_TYPE_GZH
)

type Weixin struct {
	ID                 int64  `xorm:"pk autoincr" json:"id"`
	WxId               string `xorm:"not null default '' varchar(128) index" json:"wxId"`
	Wechat             string `xorm:"not null default '' varchar(128) unique" json:"wechat"`
	NickName           string `xorm:"not null default '' varchar(256)" json:"nickName"`
	IfExecDefaultTask  int64  `xorm:"not null default 0 int" json:"ifExecDefaultTask"`
	LastHeartbeat      int64  `xorm:"not null default 0 int" json:"lastHeartbeat"`
	LastSyncContacts   int64  `xorm:"not null default 0 int" json:"lastSyncContacts"`
	TodayAddContactNum int64  `xorm:"not null default 0 int" json:"todayAddContactNum"`
	LastAddContactTime int64  `xorm:"not null default 0 int" json:"-"`
	WxType             int64  `xorm:"not null default 0 int index" json:"wxType"`
	QrcodeUrl          string `xorm:"not null default '' varchar(128)" json:"qrcodeUrl"`
	Status             int64  `xorm:"not null default 0 int index" json:"status"`
	Desc               string `xorm:"not null default '' varchar(128)" json:"desc"`
	IfWatch            int64  `xorm:"not null default 0 int index" json:"ifWatch"`
	CreatedAt          int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt          int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixin(info *Weixin) error {
	if info.Wechat == "" {
		return fmt.Errorf("wechat cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin error: %v", err)
		return err
	}
	holmes.Info("create robot weixin[%v] success.", info)

	return nil
}

func DelWeixin(info *Weixin) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin error: %v", err)
		return err
	}

	return nil
}

func GetWeixin(info *Weixin) (bool, error) {
	has, err := x.Where("wechat = ?", info.Wechat).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin from wechat[%s]", info.Wechat)
		return false, nil
	}
	return true, nil
}

func GetWeixinFromWxid(info *Weixin) (bool, error) {
	has, err := x.Where("wx_id = ?", info.WxId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin from wxid[%s]", info.WxId)
		return false, nil
	}
	return true, nil
}

func UpdateWeixinWxid(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("wx_id", "nick_name", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update wxid nickname error")
	}
	return err
}

func UpdateWeixinAddContact(info *Weixin) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	info.LastAddContactTime = now
	affected, err := x.ID(info.ID).Cols("today_add_contact_num", "last_add_contact_time", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update add contact error")
	}
	return err
}

func UpdateWeixinDesc(info *Weixin) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	affected, err := x.ID(info.ID).Cols("desc", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update desc error")
	}
	return err
}

func UpdateWeixinIfExecDefaultTask(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("if_exec_default_task", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update if_exec_default_task error")
	}
	return err
}

func UpdateWeixinLastHeartbeat(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("last_heartbeat", "updated_at").Update(info)
	//if affected == 0 {
	//	return fmt.Errorf("weixin update last_heartbeat error")
	//}
	return err
}

func UpdateWeixinLastSyncContacts(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("last_sync_contacts", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update last_sync_contacts error")
	}
	return err
}

func UpdateWeixinIfWatch(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("if_watch", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update if_watch error")
	}
	return err
}

func UpdateWeixinQrcode(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("qrcode_url", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update qrcode url error")
	}
	return err
}

func UpdateWeixinStatus(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("status", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update qrcode url error")
	}
	return err
}

func GetWeixinCount() (int64, error) {
	count, err := x.Count(&Weixin{})
	if err != nil {
		holmes.Error("get weixin list count error: %v", err)
		return 0, err
	}
	return count, nil
}

func GetWeixinList(offset, num int64) ([]Weixin, error) {
	var list []Weixin
	err := x.Limit(int(num), int(offset)).Find(&list)
	if err != nil {
		holmes.Error("get weixin list error: %v", err)
		return nil, err
	}
	return list, nil
}

func GetResourceListFromType(wxType int64) ([]Weixin, error) {
	var list []Weixin
	err := x.Where("wx_type = ?", wxType).Find(&list)
	if err != nil {
		holmes.Error("get resource list error: %v", err)
		return nil, err
	}
	return list, nil
}
