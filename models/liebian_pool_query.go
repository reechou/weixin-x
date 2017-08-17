package models

type LiebianWeixinPool struct {
	LiebianPool `xorm:"extends" json:"liebianPool"`
	Weixin      `xorm:"extends" json:"weixin"`
}

func (LiebianWeixinPool) TableName() string {
	return "liebian_pool"
}

func GetLiebianPoolFromWxId(wxid string) (*LiebianWeixinPool, error) {
	info := new(LiebianWeixinPool)
	has, err := x.Join("LEFT", "weixin", "liebian_pool.weixin_id = weixin.id").
		Where("weixin.wx_id = ?", wxid).
		Get(info)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return info, nil
}

func GetLiebianPoolWeixinList(liebianType int64) ([]LiebianWeixinPool, error) {
	list := make([]LiebianWeixinPool, 0)
	var err error
	err = x.Join("LEFT", "weixin", "liebian_pool.weixin_id = weixin.id").
		Where("liebian_pool.liebian_type = ?", liebianType).
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
