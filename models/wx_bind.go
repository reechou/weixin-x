package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type WeixinContactBindCard struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	WxId      string `xorm:"not null default '' varchar(128) unique" json:"wxId"`
	CardGid   string `xorm:"not null default '' varchar(64)" json:"cardGid"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinContactBindCard(info *WeixinContactBindCard) error {
	if info.WxId == "" {
		return fmt.Errorf("wechat cannot be nil.")
	}
	
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create weixin contact bind card error: %v", err)
		return err
	}
	holmes.Info("create weixin contact bind card[%v] success.", info)
	
	return nil
}

func GetWeixinContactBindCard(info *WeixinContactBindCard) (bool, error) {
	has, err := x.Where("wx_id = ?", info.WxId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin contact bind card from wxid[%s]", info.WxId)
		return false, nil
	}
	return true, nil
}

func UpdateWeixinContactBindCard(info *WeixinContactBindCard) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("card_gid", "updated_at").Update(info)
	return err
}
