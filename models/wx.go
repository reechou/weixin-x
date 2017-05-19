package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type Weixin struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	WxId      string `xorm:"not null default '' varchar(128) unique" json:"wxId"`
	Wechat    string `xorm:"not null default '' varchar(128) unique" json:"wechat"`
	NickName  string `xorm:"not null default '' varchar(256)" json:"nickName"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
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

func UpdateWeixinWxid(info *Weixin) error {
	info.UpdatedAt = time.Now().Unix()
	affected, err := x.ID(info.ID).Cols("wx_id", "nick_name", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("weixin update wxid nickname error")
	}
	return err
}
