package models

type WxTagFriendList struct {
	WxTagFriend   `xorm:"extends" json:"tag"`
	WeixinContact `xorm:"extends" json:"contact"`
}

func (WxTagFriendList) TableName() string {
	return "wx_tag_friend"
}

func GetWxTagFriendList(weixinId, tagId int64) ([]WxTagFriendList, error) {
	list := make([]WxTagFriendList, 0)
	err := x.Join("LEFT", "weixin_contact", "wx_tag_friend.wx_contact_id = weixin_contact.id").
		Where("wx_tag_friend.weixin_id = ?", weixinId).
		And("wx_tag_friend.tag_id = ?", tagId).
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
