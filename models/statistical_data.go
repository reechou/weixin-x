package models

import (
	"github.com/reechou/holmes"
)

const (
	S_DATA_ADD_CONTACT = iota
	S_DATA_SCREENSHOT
	S_DATA_LIEBIAN_PV
	S_DATA_LIEBIAN_UV
)

type StatisticalData struct {
	ID          int64 `xorm:"pk autoincr" json:"id"`
	TypeId      int64 `xorm:"not null default 0 int unique(ts)" json:"typeId"`
	Data        int64 `xorm:"not null default 0 int" json:"data"`
	TimeSeries  int64 `xorm:"not null default 0 int unique(ts)" json:"timeSeries"`
	LiebianType int64 `xorm:"not null default 0 int unique(ts)" json:"liebianType"`
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
	has, err := x.Where("type_id = ?", info.TypeId).
		And("time_series = ?", info.TimeSeries).
		And("liebian_type = ?", info.LiebianType).
		Get(info)
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
		And("liebian_type = ?", info.LiebianType).
		Cols("data").
		Update(info)
	return affected, err
}

type StatisticalDataMerge struct {
	TypeId     int64 `xorm:"not null default 0 int unique(ts)" json:"typeId"`
	Data       int64 `xorm:"not null default 0 int" json:"data"`
	TimeSeries int64 `xorm:"not null default 0 int unique(ts)" json:"timeSeries"`
}

func GetStatisticalDataListAll(typeId, startTime, endTime int64) ([]StatisticalDataMerge, error) {
	var list []StatisticalDataMerge
	err := x.Table("statistical_data").Where("type_id = ?", typeId).
		And("time_series >= ?", startTime).
		And("time_series <= ?", endTime).
		GroupBy("time_series").
		Cols("type_id", "sum(data)", "time_series").
		Find(&list)
	if err != nil {
		holmes.Error("get time series list error: %v", err)
		return nil, err
	}
	return list, nil
}

func GetStatisticalDataList(typeId, liebianType, startTime, endTime int64) ([]StatisticalData, error) {
	var list []StatisticalData
	var err error
	if liebianType == 0 {
		err = x.Where("type_id = ?", typeId).
			And("time_series >= ?", startTime).
			And("time_series <= ?", endTime).
			Find(&list)
	} else {
		err = x.Where("type_id = ?", typeId).
			And("liebian_type = ?", liebianType).
			And("time_series >= ?", startTime).
			And("time_series <= ?", endTime).
			Find(&list)
	}
	if err != nil {
		holmes.Error("get time series list error: %v", err)
		return nil, err
	}
	return list, nil
}
