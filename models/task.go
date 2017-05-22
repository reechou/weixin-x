package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type WeixinTask struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	TaskType  int64  `xorm:"not null default 0 int index" json:"taskType"`
	Data      string `xorm:"not null default '' varchar(1024)" json:"data"`
	IfDefault int64  `xorm:"not null default 0 int index" json:"ifDefault"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

type WeixinTaskList struct {
	ID           int64 `xorm:"pk autoincr" json:"id"`
	WeixinId     int64 `xorm:"not null default 0 int index" json:"weixinId"`
	WeixinTaskId int64 `xorm:"not null default 0 int" json:"weixinTaskId"`
	IfExec       int64 `xorm:"not null default 0 int index" json:"ifExec"`
	CreatedAt    int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt    int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateWeixinTask(info *WeixinTask) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create weixin task error: %v", err)
		return err
	}
	holmes.Info("create weixin task[%v] success.", info)

	return nil
}

func GetWeixinDefaultTaskList() ([]WeixinTask, error) {
	var list []WeixinTask
	err := x.Where("if_default = 1").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetWeixinTaskList(weixinId int64) ([]WeixinTask, error) {
	var list []WeixinTask
	err := x.Sql("select * from weixin_task where id in (select weixin_task_id from weixin_task_list where weixin_id = ? and if_exec = 0)", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateWeixinTaskList(ids []int64) error {
	_, err := x.In("id", ids).Cols("if_exec").Update(&WeixinTaskList{IfExec: 1})
	if err != nil {
		return err
	}
	return nil
}

func CreateWeixinTaskList(info *WeixinTaskList) error {
	if info.WeixinId == 0 {
		return fmt.Errorf("wechat id cannot be nil.")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create robot weixin task list error: %v", err)
		return err
	}
	holmes.Info("create robot weixin task list[%v] success.", info)

	return nil
}
