package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type WeixinVerifySetting struct {
	ID             int64  `xorm:"pk autoincr" json:"id"`
	WeixinId       int64  `xorm:"not null default 0 int index" json:"weixinId"`
	IfAutoVerified bool   `xorm:"not null default true bool" json:"ifAutoVerified"`
	Reply          string `xorm:"not null default '' varchar(1024)" json:"reply"`
	Interval       int64  `xorm:"not null default 0 int" json:"interval"`
	CreatedAt      int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt      int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinVerifySetting(info *WeixinVerifySetting) error {
	if info.WeixinId == 0 {
		return fmt.Errorf("wechat id cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin verify setting error: %v", err)
		return err
	}
	holmes.Info("create robot weixin verify setting[%v] success.", info)

	return nil
}

func DelWeixinVerifySetting(info *WeixinVerifySetting) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin verify setting error: %v", err)
		return err
	}

	return nil
}

func GetWeixinVerifySetting(info *WeixinVerifySetting) (bool, error) {
	has, err := x.Where("weixin_id = ?", info.WeixinId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin verify setting from weixin_id[%s]", info.WeixinId)
		return false, nil
	}
	return true, nil
}

func GetWeixinVerifySettingList(weixinId int64) ([]WeixinVerifySetting, error) {
	var list []WeixinVerifySetting
	err := x.Where("weixin_id = ?", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

type WeixinKeywordSetting struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	WeixinId  int64  `xorm:"not null default 0 int index" json:"weixinId"`
	ChatType  string `xorm:"not null default '' varchar(64)" json:"chatType"`
	MsgType   int64  `xorm:"not null default 0 int" json:"msgType"`
	Keyword   string `xorm:"not null default '' varchar(128)" json:"keyword"`
	Reply     string `xorm:"not null default '' varchar(1024)" json:"reply"`
	Interval  int64  `xorm:"not null default 0 int" json:"interval"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinKeywordSetting(info *WeixinKeywordSetting) error {
	if info.WeixinId == 0 {
		return fmt.Errorf("wechat id cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin keyword setting error: %v", err)
		return err
	}
	holmes.Info("create robot weixin keyword setting[%v] success.", info)

	return nil
}

func DelWeixinKeywordSetting(info *WeixinKeywordSetting) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin keyword setting error: %v", err)
		return err
	}

	return nil
}

func GetWeixinKeywordSetting(weixinId int64) ([]WeixinKeywordSetting, error) {
	var list []WeixinKeywordSetting
	err := x.Where("weixin_id = ?", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
