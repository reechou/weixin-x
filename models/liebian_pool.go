package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type LiebianType struct {
	ID           int64  `xorm:"pk autoincr" json:"id"`
	LiebianType  int64  `xorm:"not null default 0 int unique" json:"liebianType"`
	Desc         string `xorm:"not null default '' varchar(128)" json:"desc"`
	LiebianLimit int64  `xorm:"not null default 0 int" json:"liebianLimit"`
	FullLimit    int64  `xorm:"not null default 0 int" json:"fullLimit"`
	CreatedAt    int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt    int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateLiebianType(info *LiebianType) error {
	if info.LiebianType == 0 {
		return fmt.Errorf("liebian type be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create liebian type error: %v", err)
		return err
	}
	holmes.Info("create liebian type[%v] success.", info)

	return nil
}

func GetLiebianTypeList() ([]LiebianType, error) {
	var list []LiebianType
	err := x.Find(&list)
	if err != nil {
		holmes.Error("get all liebian type list error: %v", err)
		return nil, err
	}
	return list, nil
}

func DelLiebianType(info *LiebianType) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del liebian type error: %v", err)
		return err
	}

	return nil
}

func UpdateLiebianTypeLimit(info *LiebianType) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("liebian_limit", "full_limit", "updated_at").Update(info)
	return err
}

type LiebianPool struct {
	ID          int64 `xorm:"pk autoincr" json:"id"`
	LiebianType int64 `xorm:"not null default 0 int index" json:"liebianType"`
	WeixinId    int64 `xorm:"not null default 0 int index" json:"weixinId"`
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

func DelLiebianPoolList(ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("del ids cannot be nil.")
	}
	_, err := x.In("id", ids).Delete(&LiebianPool{})
	if err != nil {
		holmes.Error("del liebian pool list error: %v", err)
		return err
	}

	return nil
}

func GetLiebianPool(info *LiebianPool) (bool, error) {
	has, err := x.Where("weixin_id = ?", info.WeixinId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find liebian pool from wechat_id[%d]", info.WeixinId)
		return false, nil
	}
	return true, nil
}

func GetLiebianPoolListFromIds(ids []int64) ([]LiebianPool, error) {
	var list []LiebianPool
	err := x.In("id", ids).Find(&list)
	if err != nil {
		holmes.Error("get liebian pool list from ids error: %v", err)
		return nil, err
	}
	return list, nil
}

type LiebianErrorMsg struct {
	ID           int64  `xorm:"pk autoincr" json:"id"`
	LiebianType  int64  `xorm:"not null default 0 int index" json:"liebianType"`
	Msg          string `xorm:"not null default '' varchar(512)" json:"msg"`
	CreatedAt    int64  `xorm:"not null default 0 int index" json:"createAt"`
}

func CreateLiebianErrorMsg(info *LiebianErrorMsg) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create liebian error msg error: %v", err)
		return err
	}
	
	return nil
}

func GetLiebianErrorMsgList(liebianType int64) ([]LiebianErrorMsg, error) {
	var list []LiebianErrorMsg
	err := x.Where("liebian_type = ?", liebianType).Desc("created_at").Limit(30).Find(&list)
	if err != nil {
		holmes.Error("get liebian type error msg list error: %v", err)
		return nil, err
	}
	return list, nil
}

type LiebianOprMsg struct {
	ID           int64  `xorm:"pk autoincr" json:"id"`
	LiebianType  int64  `xorm:"not null default 0 int index" json:"liebianType"`
	Msg          string `xorm:"not null default '' varchar(512)" json:"msg"`
	CreatedAt    int64  `xorm:"not null default 0 int index" json:"createAt"`
}

func CreateLiebianOprMsg(info *LiebianOprMsg) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create liebian opr msg error: %v", err)
		return err
	}
	
	return nil
}

func GetLiebianOprMsgList(liebianType int64) ([]LiebianOprMsg, error) {
	var list []LiebianOprMsg
	err := x.Where("liebian_type = ?", liebianType).Desc("created_at").Limit(30).Find(&list)
	if err != nil {
		holmes.Error("get liebian type opr msg list error: %v", err)
		return nil, err
	}
	return list, nil
}
