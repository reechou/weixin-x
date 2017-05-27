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

func DelWeixinTask(info *WeixinTask) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del weixin task error: %v", err)
		return err
	}

	return nil
}

func UpdateWeixinTask(info *WeixinTask) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.ID).Cols("task_type", "data", "if_default").Update(info)
	return err
}

func GetWeixinTaskFromId(info *WeixinTask) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find weixin task from id[%d]", info.ID)
		return false, nil
	}
	return true, nil
}

func GetAllTaskList() ([]WeixinTask, error) {
	var list []WeixinTask
	err := x.Find(&list)
	if err != nil {
		holmes.Error("get all weixin task list error: %v", err)
		return nil, err
	}
	return list, nil
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

func UpdateWeixinTaskListFromWeixinId(weixinId int64) error {
	_, err := x.Where("weixin_id = ?", weixinId).And("if_exec = 0").Cols("if_exec").Update(&WeixinTaskList{IfExec: 1})
	if err != nil {
		return err
	}
	return nil
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

func CreateWeixinTaskInfoList(list []WeixinTaskList) error {
	if len(list) == 0 {
		return nil
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create weixin task list error: %v", err)
		return err
	}
	return nil
}
