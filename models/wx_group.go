package models

import (
	"fmt"
	"time"
	
	"github.com/reechou/holmes"
)

type WeixinGroup struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	GroupName string `xorm:"not null default '' varchar(128)" json:"groupName"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

type WeixinGroupMember struct {
	ID        int64 `xorm:"pk autoincr" json:"id"`
	GroupId   int64 `xorm:"not null default 0 int index" json:"groupId"`
	WeixinId  int64 `xorm:"not null default 0 int" json:"weixinId"`
	CreatedAt int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinGroup(info *WeixinGroup) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create weixin group error: %v", err)
		return err
	}
	holmes.Info("create weixin group[%v] success.", info)
	
	return nil
}

func GetWeixinGroupList() ([]WeixinGroup, error) {
	var list []WeixinGroup
	err := x.Find(&list)
	if err != nil {
		holmes.Error("get all weixin group list error: %v", err)
		return nil, err
	}
	return list, nil
}

func DelWeixinGroup(info *WeixinGroup) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin group error: %v", err)
		return err
	}
	
	return nil
}

func CreateWeixinGroupMemberList(list []WeixinGroupMember) error {
	if len(list) == 0 {
		return nil
	}
	for i := 0; i < len(list); i++ {
		list[i].CreatedAt = time.Now().Unix()
		list[i].UpdatedAt = time.Now().Unix()
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create weixin group member list error: %v", err)
		return err
	}
	return nil
}

func DelWeixinGroupMember(info *WeixinGroupMember) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin group member error: %v", err)
		return err
	}
	
	return nil
}

func GetWeixinGroupMemberList(groupId int64) ([]WeixinGroupMember, error) {
	var list []WeixinGroupMember
	err := x.Where("group_id = ?", groupId).Find(&list)
	if err != nil {
		holmes.Error("get group member list error: %v", err)
		return nil, err
	}
	return list, nil
}
