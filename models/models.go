package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
)

var x *xorm.Engine

func InitDB(cfg *config.Config) {
	var err error
	x, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		cfg.DBInfo.User,
		cfg.DBInfo.Pass,
		cfg.DBInfo.Host,
		cfg.DBInfo.DBName))
	if err != nil {
		holmes.Fatal("Fail to init new engine: %v", err)
	}
	x.SetMapper(core.GonicMapper{})
	x.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	// if need show raw sql in log
	if cfg.IfShowSql {
		x.ShowSQL(true)
	}

	// sync tables
	if err = x.Sync2(new(Weixin),
		new(WeixinVerifySetting),
		new(WeixinKeywordSetting),
		new(WeixinVerify),
		new(WeixinKeyword),
		new(WeixinTask),
		new(WeixinTaskList),
		new(WeixinContact),
		new(WeixinContactBindCard),
		new(WxFriendTag),
		new(WxTagFriend),
		new(TimerTask),
		new(LiebianType),
		new(WeixinGroup),
		new(WeixinGroupMember),
		new(LiebianPool),
		new(QrcodeBind),
		new(StatisticalData),
		new(LiebianErrorMsg),
		new(LiebianOprMsg),
		new(WeixinChatroomSetting),
		new(WeixinChatroomSettingDetail)); err != nil {
		holmes.Fatal("Fail to sync database: %v", err)
	}
}
