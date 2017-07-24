package models

import (
	"time"
	
	"github.com/reechou/holmes"
)

type LiebianPool struct {
	ID          int64 `xorm:"pk autoincr" json:"id"`
	LiebianType int64 `xorm:"not null default 0 int" json:"liebianType"`
	WeixinId    int64 `xorm:"not null default 0 int" json:"weixinId"`
	CreatedAt   int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt   int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateLiebianPoolList(list []LiebianPool) error {
	if len(list) == 0 {
		return nil
	}
	for i := 0; i < len(list); i++ {
		list[i].CreatedAt = time.Now().Unix()
		list[i].UpdatedAt = time.Now().Unix()
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create weixin liebian pool list error: %v", err)
		return err
	}
	return nil
}
