package controller

import (
	"encoding/json"
	"net/http"
	
	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/proto"
)

func (self *Logic) CreateWeixinChatroomSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.ChatroomCommonSetting{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateChatroomSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	setting, err := json.Marshal(req)
	if err != nil {
		holmes.Error("json marshal reply error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	chatroomSetting := &models.WeixinChatroomSetting{
		Setting: string(setting),
	}
	err = models.CreateWeixinChatroomSetting(chatroomSetting)
	if err != nil {
		holmes.Error("create weixin chatroom setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) DeleteWeixinChatroomSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &models.WeixinChatroomSetting{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteWeixinChatroomSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err := models.DelWeixinChatroomSetting(req)
	if err != nil {
		holmes.Error("delete weixin chatroom setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) UpdateWeixinChatroomSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &models.WeixinChatroomSetting{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("UpdateWeixinChatroomSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err := models.UpdateWeixinChatroomSetting(req)
	if err != nil {
		holmes.Error("update weixin chatroom setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetAllChatroomSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	list, err := models.GetAllWeixinChatroomSettingList()
	if err != nil {
		holmes.Error("get all weixin chatroom setting list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) CreateWeixinChatroomSettingDetailList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	var req []models.WeixinChatroomSettingDetail
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		holmes.Error("CreateWeixinChatroomSettingDetailList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err := models.CreateWeixinChatroomSettingDetailList(req)
	if err != nil {
		holmes.Error("create weixin chatroom setting detail list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) DeleteWeixinChatroomSettingDetail(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &models.WeixinChatroomSettingDetail{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteWeixinChatroomSettingDetail json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err := models.DelWeixinChatroomSettingDetail(req)
	if err != nil {
		holmes.Error("delete weixin chatroom setting detail error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetWeixinChatroomSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.Weixin{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetWeixinChatroomSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	holmes.Debug("get weixin chatroom setting req: %v", req)
	
	chatroomSetting, err := models.GetWxChatroomSetting(req.WxId)
	if err != nil {
		holmes.Error("get wx chatroom setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	if chatroomSetting.WeixinChatroomSettingDetail.ID == 0 {
		return
	}

	var setting proto.ChatroomCommonSetting
	err = json.Unmarshal([]byte(chatroomSetting.WeixinChatroomSetting.Setting), &setting)
	if err != nil {
		holmes.Error("json unmarshal chatroom setting reply error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	setting.ChatroomRole = chatroomSetting.WeixinChatroomSettingDetail.ChatroomRole
	rsp.Data = setting
}
