package models

import (
	"time"

	"github.com/reechou/holmes"
)

type WeixinContact struct {
	ID             int64  `xorm:"pk autoincr" json:"id"`
	WeixinId       int64  `xorm:"not null default 0 int index" json:"weixinId"`
	UserName       string `xorm:"not null default '' varchar(128) unique" json:"userName"`
	AliasName      string `xorm:"not null default '' varchar(128)" json:"aliasName"`
	NickName       string `xorm:"not null default '' varchar(128)" json:"nickName"`
	PhoneNumber    string `xorm:"not null default '' varchar(128)" json:"phoneNumber"`
	Country        string `xorm:"not null default '' varchar(128)" json:"country"`
	Province       string `xorm:"not null default '' varchar(128)" json:"province"`
	City           string `xorm:"not null default '' varchar(128)" json:"city"`
	Sex            int64  `xorm:"not null default 0 int" json:"sex"`
	Remark         string `xorm:"not null default '' varchar(128)" json:"remark"`
	AddContactTime int64  `xorm:"not null default 0 int index" json:"addContactTime"`
	CreatedAt      int64  `xorm:"not null default 0 int" json:"createAt"`
}

func CreateWeixinContact(info *WeixinContact) error {
	info.CreatedAt = time.Now().Unix()
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create weixin contact error: %v", err)
		return err
	}

	return nil
}

func GetWeixinContact(info *WeixinContact) (bool, error) {
	has, err := x.Where("user_name = ?", info.UserName).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		//holmes.Debug("cannot find weixin contact from username[%s]", info.UserName)
		return false, nil
	}
	return true, nil
}

func GetWeixinContactCount(weixinId int64) (int64, error) {
	count, err := x.Where("weixin_id = ?", weixinId).Count(&WeixinContact{})
	if err != nil {
		holmes.Error("get weixin contact count error: %v", err)
		return 0, err
	}
	return count, nil
}

func GetWeixinContactList(weixinId, offset, num int64) ([]WeixinContact, error) {
	var list []WeixinContact
	err := x.Where("weixin_id = ?", weixinId).
		Limit(int(num), int(offset)).
		Desc("add_contact_time").
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetWeixinContactListFromTime(weixinId, startTime, endTime int64) ([]WeixinContact, error) {
	var list []WeixinContact
	err := x.Where("weixin_id = ?", weixinId).
		And("add_contact_time >= ?", startTime).
		And("add_contact_time <= ?", endTime).
		Desc("add_contact_time").
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
