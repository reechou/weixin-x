package models

type WxFriendTag struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	TagType   int64  `xorm:"not null default 0 int" json:"tagType"`
	GroupName string `xorm:"not null default '' varchar(128)" json:"groupName"`
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

func GetWxFriendTagList() ([]WxFriendTag, error) {
	var list []WxFriendTag
	err := x.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
