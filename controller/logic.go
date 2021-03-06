package controller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/ext"
)

type Logic struct {
	sync.Mutex

	cw  *ContactWorker
	ttw *TimerTaskWorker
	wmm *WeixinMonitorManager
	dsw *DataStatisticsWorker
	
	alarm *ext.AlarmExt

	cfg *config.Config
}

func NewLogic(cfg *config.Config) *Logic {
	l := &Logic{
		cfg: cfg,
	}
	l.cw = NewContactWorker()
	l.ttw = NewTimerTaskWorker(cfg)
	l.init()

	models.InitDB(cfg)

	return l
}

func (self *Logic) init() {
	http.HandleFunc("/robot/receive_msg", self.RobotReceiveMsg)

	http.HandleFunc("/weixin/create_resource", self.CreateResource)
	http.HandleFunc("/weixin/get_resource_pool", self.GetResourcePool)
	http.HandleFunc("/weixin/create_weixin", self.CreateWeixin)
	http.HandleFunc("/weixin/update_weixin_desc", self.UpdateWeixinDesc)
	http.HandleFunc("/weixin/update_weixin_qrcode", self.UpdateWeixinQrcode)
	http.HandleFunc("/weixin/update_weixin_status", self.UpdateWeixinStatus)
	http.HandleFunc("/weixin/delete_weixin", self.DeleteWeixin)
	http.HandleFunc("/weixin/create_verify_setting", self.CreateWeixinVerifySetting)
	http.HandleFunc("/weixin/get_verify_setting", self.GetWeixinVerifySetting)
	http.HandleFunc("/weixin/delete_verify_setting", self.DeleteWeixinVerifySetting)
	http.HandleFunc("/weixin/update_verify_setting", self.UpdateWeixinVerifySetting)
	http.HandleFunc("/weixin/create_verify", self.CreateWeixinVerify)
	http.HandleFunc("/weixin/delete_verify", self.DeleteWeixinVerify)
	http.HandleFunc("/weixin/create_keyword_setting", self.CreateWeixinKeywordSetting)
	http.HandleFunc("/weixin/get_keyword_setting", self.GetWeixinKeywordSetting)
	http.HandleFunc("/weixin/delete_keyword_setting", self.DeleteWeixinKeywordSetting)
	http.HandleFunc("/weixin/update_keyword_setting", self.UpdateWeixinKeywordSetting)
	http.HandleFunc("/weixin/create_keyword", self.CreateWeixinKeyword)
	http.HandleFunc("/weixin/delete_keyword", self.DeleteWeixinKeyword)
	http.HandleFunc("/weixin/get_setting", self.GetWeixinSetting)
	http.HandleFunc("/weixin/get_setting_from_id", self.GetWeixinSettingFromId)
	http.HandleFunc("/weixin/create_task", self.CreateTask)
	http.HandleFunc("/weixin/delete_task", self.DeleteTask)
	http.HandleFunc("/weixin/update_task", self.UpdateTask)
	http.HandleFunc("/weixin/get_task_from_id", self.GetTaskFromId)
	http.HandleFunc("/weixin/create_weixin_task", self.CreateWeixinTask)
	http.HandleFunc("/weixin/batch_create_weixin_task", self.BatchCreateTaskList)
	http.HandleFunc("/weixin/get_task", self.GetTask)
	http.HandleFunc("/weixin/sync_contacts", self.SyncContacts)
	http.HandleFunc("/weixin/add_contact", self.AddContact)
	http.HandleFunc("/weixin/get_contact_bind", self.GetWeixinContactBind)
	http.HandleFunc("/weixin/get_weixin_friends", self.GetWeixinFriends)
	http.HandleFunc("/weixin/get_weixin_friends_tags", self.GetWxFriendTagList)
	http.HandleFunc("/weixin/get_weixin_friends_from_time", self.GetWeixinFriendsFromTime)
	http.HandleFunc("/weixin/get_weixin_friends_from_tag", self.GetWeixinFriendsFromTag)
	http.HandleFunc("/weixin/delete_weixin_friend_tag", self.DeleteWxFriendTag)
	http.HandleFunc("/weixin/create_selected_friends_task", self.CreateSelectedFriendsTask)
	http.HandleFunc("/weixin/batch_create_selected_friends_task", self.BatchCreateSelectedFriendsTask)
	http.HandleFunc("/weixin/create_timer_task", self.CreateTimerTask)
	http.HandleFunc("/weixin/get_timer_task_list", self.GetTimerTaskList)
	http.HandleFunc("/weixin/delete_timer_task", self.DeleteTimerTask)

	// liebian
	http.HandleFunc("/weixin/create_liebian_type", self.CreateLiebianType)
	http.HandleFunc("/weixin/delete_liebian_type", self.DeleteLiebianType)
	http.HandleFunc("/weixin/update_liebian_type_limit", self.UpdateLiebianTypeLimit)
	http.HandleFunc("/weixin/get_liebian_type_list", self.GetLiebianTypeList)
	http.HandleFunc("/weixin/create_weixin_group", self.CreateWeixinGroup)
	http.HandleFunc("/weixin/delete_weixin_group", self.DeleteWeixinGroup)
	http.HandleFunc("/weixin/get_weixin_group_list", self.GetWeixinGroupList)
	http.HandleFunc("/weixin/create_weixin_group_member_list", self.CreateWeixinGroupMemberList)
	http.HandleFunc("/weixin/delete_weixin_group_member", self.DeleteWeixinGroupMember)
	http.HandleFunc("/weixin/get_weixin_group_member_list", self.GetWeixinGroupMemberList)
	http.HandleFunc("/weixin/create_lianbian_pool", self.CreateLiebianPool)
	http.HandleFunc("/weixin/delete_lianbian_pool", self.DeleteLiebianPool)
	http.HandleFunc("/weixin/get_lianbian_pool", self.GetLiebianPool)
	http.HandleFunc("/weixin/get_user_lianbian_pool", self.GetUserLiebianInfo)
	http.HandleFunc("/weixin/get_liebian_error_msg", self.GetLiebianErrorMsgList)
	http.HandleFunc("/weixin/get_liebian_opr_msg", self.GetLiebianOprMsgList)
	
	// chatroom setting
	http.HandleFunc("/weixin/create_chatroom_setting", self.CreateWeixinChatroomSetting)
	http.HandleFunc("/weixin/delete_chatroom_setting", self.DeleteWeixinChatroomSetting)
	http.HandleFunc("/weixin/update_chatroom_setting", self.UpdateWeixinChatroomSetting)
	http.HandleFunc("/weixin/get_all_chatroom_setting", self.GetAllChatroomSetting)
	http.HandleFunc("/weixin/batch_create_chatroom_setting_detail", self.CreateWeixinChatroomSettingDetailList)
	http.HandleFunc("/weixin/delete_chatroom_setting_detail", self.DeleteWeixinChatroomSettingDetail)
	http.HandleFunc("/weixin/get_chatroom_setting", self.GetWeixinChatroomSetting)
	http.HandleFunc("/weixin/get_chatroom_setting_from_weixin_id", self.GetWeixinChatroomSettingFromWeixinId)
	
	// monitor
	http.HandleFunc("/monitor/get_data", self.GetDataStatistical)

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
	// start monitor
	self.alarm = ext.NewAlarmExt(self.cfg)
	self.wmm = NewWeixinMonitorManager(self.alarm, self.cfg)
	self.dsw = NewDataStatisticsWorker(self.cfg)

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
