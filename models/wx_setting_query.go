package models

type WeixinVerifySettingDetail struct {
	WeixinVerify        `xorm:"extends" json:"weixinVerify"`
	WeixinVerifySetting `xorm:"extends" json:"weixinVerifySetting"`
	Weixin              `xorm:"extends" json:"weixin"`
}

func (WeixinVerifySettingDetail) TableName() string {
	return "weixin_verify"
}

func GetWeixinVerifySettingFromWeixin(weixinId int64) (*WeixinVerifySettingDetail, error) {
	detail := new(WeixinVerifySettingDetail)
	var err error
	_, err = x.Join("LEFT", "weixin_verify_setting", "weixin_verify.weixin_verify_setting_id = weixin_verify_setting.id").
		Join("LEFT", "weixin", "weixin_verify.weixin_id = weixin.id").
		Where("weixin_verify.weixin_id = ?", weixinId).
		Get(detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

type WeixinKeywordSettingList struct {
	WeixinKeyword        `xorm:"extends" json:"weixinKeyword"`
	WeixinKeywordSetting `xorm:"extends" json:"weixinKeywordSetting"`
	Weixin               `xorm:"extends" json:"weixin"`
}

func (WeixinKeywordSettingList) TableName() string {
	return "weixin_keyword"
}

func GetWeixinKeywordSettingListFromWeixin(weixinId int64) ([]WeixinKeywordSettingList, error) {
	list := make([]WeixinKeywordSettingList, 0)
	var err error
	err = x.Join("LEFT", "weixin_keyword_setting", "weixin_keyword.weixin_keyword_setting_id = weixin_keyword_setting.id").
		Join("LEFT", "weixin", "weixin_keyword.weixin_id = weixin.id").
		Where("weixin_keyword.weixin_id = ?", weixinId).
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

type WxChatroomSetting struct {
	WeixinChatroomSettingDetail `xorm:"extends" json:"weixinChatroomSettingDetail"`
	WeixinChatroomSetting       `xorm:"extends" json:"weixinChatroomSetting"`
	Weixin                      `xorm:"extends" json:"weixin"`
}

func (WxChatroomSetting) TableName() string {
	return "weixin_chatroom_setting_detail"
}

func GetWxChatroomSetting(wxid string) (*WxChatroomSetting, error) {
	detail := new(WxChatroomSetting)
	var err error
	_, err = x.Join("LEFT", "weixin_chatroom_setting", "weixin_chatroom_setting_detail.weixin_chatroom_setting_id = weixin_chatroom_setting.id").
		Join("LEFT", "weixin", "weixin_chatroom_setting_detail.weixin_id = weixin.id").
		Where("weixin.wx_id = ?", wxid).
		Get(detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func GetWxChatroomSettingFromWeixin(weixinId int64) (*WxChatroomSetting, error) {
	detail := new(WxChatroomSetting)
	var err error
	_, err = x.Join("LEFT", "weixin_chatroom_setting", "weixin_chatroom_setting_detail.weixin_chatroom_setting_id = weixin_chatroom_setting.id").
		Join("LEFT", "weixin", "weixin_chatroom_setting_detail.weixin_id = weixin.id").
		Where("weixin_chatroom_setting_detail.weixin_id = ?", weixinId).
		Get(detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}
