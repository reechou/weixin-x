package models

import (
	"time"
	
	"github.com/reechou/holmes"
)

type QrcodeBind struct {
	ID          int64  `xorm:"pk autoincr" json:"id"`
	AppId       string `xorm:"not null default '' varchar(64) unique(user_qrcode_bind)" json:"appId"`
	OpenId      string `xorm:"not null default '' varchar(128) unique(user_qrcode_bind)" json:"appId"`
	LiebianType int64  `xorm:"not null default 0 int unique(user_qrcode_bind)" json:"liebianType"`
	BindQrcode  string `xorm:"not null default '' varchar(256)" json:"bindQrcode"`
	CreatedAt   int64  `xorm:"not null default 0 int" json:"createAt"`
}

func CreateQrcodeBind(info *QrcodeBind) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create qrcode bind error: %v", err)
		return err
	}
	holmes.Info("create qrcode bind[%v] success.", info)
	
	return nil
}

func GetQrcodeBind(info *QrcodeBind) (bool, error) {
	has, err := x.Where("app_id = ?", info.AppId).
		And("open_id = ?", info.OpenId).
		And("liebian_type = ?", info.LiebianType).
		Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}
