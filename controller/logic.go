package controller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
	"github.com/reechou/weixin-x/models"
)

type Logic struct {
	sync.Mutex

	cfg *config.Config
}

func NewLogic(cfg *config.Config) *Logic {
	l := &Logic{
		cfg: cfg,
	}
	l.init()

	models.InitDB(cfg)

	return l
}

func (self *Logic) init() {
	http.HandleFunc("/robot/receive_msg", self.RobotReceiveMsg)

	http.HandleFunc("/weixin/create_weixin", self.CreateWeixin)
	http.HandleFunc("/weixin/delete_weixin", self.DeleteWeixin)
	http.HandleFunc("/weixin/create_verify_setting", self.CreateWeixinVerifySetting)
	http.HandleFunc("/weixin/delete_verify_setting", self.DeleteWeixinVerifySetting)
	http.HandleFunc("/weixin/create_verify", self.CreateWeixinVerify)
	http.HandleFunc("/weixin/delete_verify", self.DeleteWeixinVerify)
	http.HandleFunc("/weixin/create_keyword_setting", self.CreateWeixinKeywordSetting)
	http.HandleFunc("/weixin/delete_keyword_setting", self.DeleteWeixinKeywordSetting)
	http.HandleFunc("/weixin/create_keyword", self.CreateWeixinKeyword)
	http.HandleFunc("/weixin/delete_keyword", self.DeleteWeixinKeyword)
	http.HandleFunc("/weixin/get_setting", self.GetWeixinSetting)
	http.HandleFunc("/weixin/get_setting_from_id", self.GetWeixinSettingFromId)
	http.HandleFunc("/weixin/create_task", self.CreateTask)
	http.HandleFunc("/weixin/delete_task", self.DeleteTask)
	http.HandleFunc("/weixin/create_weixin_task", self.CreateWeixinTask)
	http.HandleFunc("/weixin/get_task", self.GetTask)
	
	http.HandleFunc("/weixin/get_all_weixin", self.GetAllWeixin)
	http.HandleFunc("/weixin/get_all_verify", self.GetAllVerifySetting)
	http.HandleFunc("/weixin/get_all_keyword", self.GetAllKeywordSetting)
	http.HandleFunc("/weixin/get_all_task", self.GetAllTask)
}

func (self *Logic) Run() {
	defer holmes.Start(holmes.LogFilePath("./log"),
		holmes.EveryDay,
		holmes.AlsoStdout,
		holmes.DebugLevel).Stop()

	if self.cfg.Debug {
		EnableDebug()
	}

	holmes.Info("server starting on[%s]..", self.cfg.Host)
	holmes.Infoln(http.ListenAndServe(self.cfg.Host, nil))
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func WriteBytes(w http.ResponseWriter, code int, v []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(v)
}

func EnableDebug() {

}
