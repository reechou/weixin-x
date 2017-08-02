package models

import (
	"github.com/reechou/holmes"
)

const (
	S_DATA_ADD_CONTACT = iota
	S_DATA_SCREENSHOT
)

type StatisticalData struct {
	ID         int64 `xorm:"pk autoincr" json:"id"`
	TypeId     int64 `xorm:"not null default 0 int unique(ts)" json:"typeId"`
	Data       int64 `xorm:"not null default 0 int" json:"data"`
	TimeSeries int64 `xorm:"not null default 0 int unique(ts)" json:"timeSeries"`
}

func CreateStatisticalData(info *StatisticalData) error {
	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create statistical data error: %v", err)
		return err
	}
	
	return nil
}

func GetStatisticalData(info *StatisticalData) (bool, error) {
	has, err := x.Where("type_id = ?", info.TypeId).And("time_series = ?", info.TimeSeries).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func UpdateStatisticalData(info *StatisticalData) (int64, error) {
	affected, err := x.Where("type_id = ?", info.TypeId).
		And("time_series = ?", info.TimeSeries).
		Cols("data").
		Update(info)
	return affected, err
}

func GetStatisticalDataList(typeId, startTime, endTime int64) ([]StatisticalData, error) {
	var list []StatisticalData
	err := x.Where("type_id = ?", typeId).
		And("time_series >= ?", startTime).
		And("time_series <= ?", endTime).
		Find(&list)
	if err != nil {
		holmes.Error("get time series list error: %v", err)
		return nil, err
	}
	return list, nil
}
