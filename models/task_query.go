package models

type WxTaskList struct {
	WeixinTaskList `xorm:"extends"`
	WeixinTask     `xorm:"extends"`
}

func (WxTaskList) TableName() string {
	return "weixin_task_list"
}

func GetWxTaskList(weixinId int64) ([]WxTaskList, error) {
	list := make([]WxTaskList, 0)
	err := x.Join("LEFT", "weixin_task", "weixin_task_list.weixin_task_id = weixin_task.id").
		Where("weixin_task_list.weixin_id = ?", weixinId).
		And("weixin_task_list.if_exec = 0").
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
