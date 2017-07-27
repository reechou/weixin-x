package models

type WeixinGroupMemberList struct {
	WeixinGroupMember `xorm:"extends" json:"weixinGroupMember"`
	Weixin            `xorm:"extends" json:"weixin"`
}

func (WeixinGroupMemberList) TableName() string {
	return "weixin_group_member"
}

func GetWeixinGroupMemberDetailList(groupId int64) ([]WeixinGroupMemberList, error) {
	list := make([]WeixinGroupMemberList, 0)
	var err error
	err = x.Join("LEFT", "weixin", "weixin_group_member.weixin_id = weixin.id").
		Where("weixin_group_member.group_id = ?", groupId).
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
