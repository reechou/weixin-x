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

type WeixinVerify struct {
	ID                    int64 `xorm:"pk autoincr" json:"id"`
	WeixinId              int64 `xorm:"not null default 0 int index" json:"weixinId"`
	WeixinVerifySettingId int64 `xorm:"not null default 0 int index" json:"weixinVerifySettingId"`
	CreatedAt             int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt             int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinVerifySetting(info *WeixinVerifySetting) error {
	//if info.WeixinId == 0 {
	//	return fmt.Errorf("wechat id cannot be nil.")
	//}

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

func UpdateWeixinVerifySetting(info *WeixinVerifySetting) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.ID).Cols("if_auto_verified", "reply", "interval").Update(info)
	return err
}

func GetWeixinVerifySetting(info *WeixinVerifySetting) (bool, error) {
	has, err := x.Where("weixin_id = ?", info.WeixinId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin verify setting from weixin_id[%d]", info.WeixinId)
		return false, nil
	}
	return true, nil
}

func GetWeixinVerifySettingDetail(weixinId int64) (*WeixinVerifySetting, error) {
	setting := new(WeixinVerifySetting)
	has, err := x.Sql("select * from weixin_verify_setting where id = (select weixin_verify_setting_id from weixin_verify where weixin_id = ? limit 1)", weixinId).Get(setting)
	if err != nil {
		return nil, err
	}
	if !has {
		holmes.Debug("cannot find weixin verify setting from weixin_id[%d]", weixinId)
		return nil, nil
	}
	return setting, nil
}

func GetWeixinVerifySettingFromId(info *WeixinVerifySetting) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin verify setting from id[%d]", info.ID)
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

func CreateWeixinVerify(info *WeixinVerify) error {
	if info.WeixinId == 0 {
		return fmt.Errorf("wechat id cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin verify error: %v", err)
		return err
	}
	holmes.Info("create robot weixin verify[%v] success.", info)

	return nil
}

func DelWeixinVerify(info *WeixinVerify) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin verify error: %v", err)
		return err
	}

	return nil
}

func GetWeixinVerify(info *WeixinVerify) (bool, error) {
	has, err := x.Where("weixin_id = ?", info.WeixinId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin verify from weixin_id[%d]", info.WeixinId)
		return false, nil
	}
	return true, nil
}

func GetAllVerifyList() ([]WeixinVerifySetting, error) {
	var list []WeixinVerifySetting
	err := x.Find(&list)
	if err != nil {
		holmes.Error("get all weixin verify setting list error: %v", err)
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

type WeixinKeyword struct {
	ID                     int64 `xorm:"pk autoincr" json:"id"`
	WeixinId               int64 `xorm:"not null default 0 int index" json:"weixinId"`
	WeixinKeywordSettingId int64 `xorm:"not null default 0 int index" json:"weixinKeywordSettingId"`
	CreatedAt              int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt              int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinKeywordSetting(info *WeixinKeywordSetting) error {
	//if info.WeixinId == 0 {
	//	return fmt.Errorf("wechat id cannot be nil.")
	//}

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

func UpdateWeixinKeywordSetting(info *WeixinKeywordSetting) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.ID).Cols("chat_type", "msg_type", "keyword", "reply", "interval").Update(info)
	return err
}

func GetWeixinKeywordSettingList(weixinId int64) ([]WeixinKeywordSetting, error) {
	var list []WeixinKeywordSetting
	err := x.Sql("select * from weixin_keyword_setting where id in (select weixin_keyword_setting_id from weixin_keyword where weixin_id = ?)", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetWeixinKeywordSetting(weixinId int64) ([]WeixinKeywordSetting, error) {
	var list []WeixinKeywordSetting
	err := x.Where("weixin_id = ?", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func CreateWeixinKeyword(info *WeixinKeyword) error {
	if info.WeixinId == 0 {
		return fmt.Errorf("wechat id cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin keyword error: %v", err)
		return err
	}
	holmes.Info("create robot weixin keyword[%v] success.", info)

	return nil
}

func CreateWeixinKeywordList(list []WeixinKeyword) error {
	if len(list) == 0 {
		return nil
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create weixin keyword list error: %v", err)
		return err
	}
	return nil
}

func DelWeixinKeyword(info *WeixinKeyword) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin keyword error: %v", err)
		return err
	}

	return nil
}

func GetWeixinKeyword(info *WeixinKeyword) (bool, error) {
	has, err := x.Where("weixin_id = ?", info.WeixinId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin keyword from weixin_id[%d]", info.WeixinId)
		return false, nil
	}
	return true, nil
}

func GetAllKeywordList() ([]WeixinKeywordSetting, error) {
	var list []WeixinKeywordSetting
	err := x.Find(&list)
	if err != nil {
		holmes.Error("get all weixin keyword list error: %v", err)
		return nil, err
	}
	return list, nil
}

func GetWeixinKeywordSettingFromId(info *WeixinKeywordSetting) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin keyword setting from id[%d]", info.ID)
		return false, nil
	}
	return true, nil
}
