package models

import (
	"fmt"
	"time"
	
	"github.com/reechou/holmes"
)

type TimerTask struct {
	ID        int64 `xorm:"pk autoincr" json:"id"`
	WeixinId  int64 `xorm:"not null default 0 int index" json:"weixinId"`
	TagId     int64 `xorm:"not null default 0 int" json:"tagId"`
	TimeId    int64 `xorm:"not null default 0 int" json:"timeId"`
	TaskId    int64 `xorm:"not null default 0 int" json:"taskId"`
	CreatedAt int64 `xorm:"not null default 0 int" json:"createAt"`
	UpdatedAt int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateTimerTask(info *TimerTask) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create timer task error: %v", err)
		return err
	}
	holmes.Info("create timer task[%v] success.", info)
	
	return nil
}

func GetTimerTaskList() ([]TimerTask, error) {
	var list []TimerTask
	err := x.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetTimerTaskListFromWeixin(weixinId int64) ([]TimerTask, error) {
	var list []TimerTask
	err := x.Where("weixin_id = ?", weixinId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func DelTimerTask(info *TimerTask) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del timer task error: %v", err)
		return err
	}
	
	return nil
}
