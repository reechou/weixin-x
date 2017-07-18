package models

type WxFriendGroup struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	GroupName string `xorm:"not null default '' varchar(128)" json:"groupName"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}
