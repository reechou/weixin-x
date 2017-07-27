package models

import (
	"time"

	"github.com/reechou/holmes"
)

type WxFriendTag struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	TagType   int64  `xorm:"not null default 0 int" json:"tagType"`
	GroupName string `xorm:"not null default '' varchar(128) unique" json:"groupName"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

type WxTagFriend struct {
	ID          int64 `xorm:"pk autoincr" json:"id"`
	WeixinId    int64 `xorm:"not null default 0 int index" json:"weixinId"`
	TagId       int64 `xorm:"not null default 0 int index" json:"tagId"`
	WxContactId int64 `xorm:"not null default 0 int index" json:"wxContactId"`
	CreatedAt   int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt   int64 `xorm:"not null default 0 int" json:"-"`
}

func GetWxFriendTag(info *WxFriendTag) (bool, error) {
	has, err := x.Where("group_name = ?", info.GroupName).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func GetWxFriendTagList() ([]WxFriendTag, error) {
	var list []WxFriendTag
	err := x.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func CreateWxTagFriend(info *WxTagFriend) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create WxTagFriend error: %v", err)
		return err
	}
	holmes.Info("create WxTagFriend[%v] success.", info)

	return nil
}

func DelWxTagFriend(info *WxTagFriend) error {
	_, err := x.Where("weixin_id = ?", info.WeixinId).
		And("tag_id = ?", info.TagId).
		And("wx_contact_id = ?", info.WxContactId).
		Delete(info)
	if err != nil {
		return err
	}
	return nil
}
