package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/proto"
)

func (self *Logic) CreateWeixin(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.Weixin{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixin json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	weixin := &models.Weixin{
		WxId:     req.WxId,
		Wechat:   req.Wechat,
		NickName: req.NickName,
	}
	err := models.CreateWeixin(weixin)
	if err != nil {
		holmes.Error("create weixin error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = weixin.ID

	if req.VerifyId != 0 {
		wv := &models.WeixinVerify{
			WeixinId:              weixin.ID,
			WeixinVerifySettingId: req.VerifyId,
		}
		err = models.CreateWeixinVerify(wv)
		if err != nil {
			holmes.Error("create weixin verify error: %v", err)
		}
	}
	if req.KeywordIds != nil {
		var wks []models.WeixinKeyword
		for _, v := range req.KeywordIds {
			wks = append(wks, models.WeixinKeyword{
				WeixinId:               weixin.ID,
				WeixinKeywordSettingId: v,
				CreatedAt:              time.Now().Unix(),
				UpdatedAt:              time.Now().Unix(),
			})
		}
		err = models.CreateWeixinKeywordList(wks)
		if err != nil {
			holmes.Error("create weixin keyword list error: %v", err)
		}
	}
	if req.TaskIds != nil {
		var wts []models.WeixinTaskList
		for _, v := range req.TaskIds {
			wts = append(wts, models.WeixinTaskList{
				WeixinId:     weixin.ID,
				WeixinTaskId: v,
				CreatedAt:    time.Now().Unix(),
				UpdatedAt:    time.Now().Unix(),
			})
		}
		err = models.CreateWeixinTaskInfoList(wts)
		if err != nil {
			holmes.Error("create weixin task list error: %v", err)
		}
	}
}

func (self *Logic) CreateWeixinVerifySetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.VerifySetting{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinVerifySetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	//weixin := &models.Weixin{
	//	Wechat: req.Wechat,
	//}
	//has, err := models.GetWeixin(weixin)
	//if err != nil {
	//	holmes.Error("get weixin error: %v", err)
	//	rsp.Code = proto.RESPONSE_ERR
	//	return
	//}
	//if !has {
	//	holmes.Error("Has no this weixin[%s]", req.Wechat)
	//	rsp.Code = proto.RESPONSE_ERR
	//	return
	//}

	reply, err := json.Marshal(req.Reply)
	if err != nil {
		holmes.Error("json marshal reply error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	setting := &models.WeixinVerifySetting{
		IfAutoVerified: req.IfAutoVerified,
		Reply:          string(reply),
		Interval:       req.Interval,
	}
	err = models.CreateWeixinVerifySetting(setting)
	if err != nil {
		holmes.Error("create weixin verify setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) CreateWeixinVerify(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinVerify{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinVerify json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateWeixinVerify(req)
	if err != nil {
		holmes.Error("create weixin verify error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) CreateWeixinKeywordSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.KeywordSetting{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinKeywordSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	//weixin := &models.Weixin{
	//	Wechat: req.Wechat,
	//}
	//has, err := models.GetWeixin(weixin)
	//if err != nil {
	//	holmes.Error("get weixin error: %v", err)
	//	rsp.Code = proto.RESPONSE_ERR
	//	return
	//}
	//if !has {
	//	holmes.Error("Has no this weixin[%s]", req.Wechat)
	//	rsp.Code = proto.RESPONSE_ERR
	//	return
	//}

	reply, err := json.Marshal(req.Reply)
	if err != nil {
		holmes.Error("json marshal reply error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	setting := &models.WeixinKeywordSetting{
		ChatType: req.ChatType,
		MsgType:  req.MsgType,
		Keyword:  req.Keyword,
		Reply:    string(reply),
		Interval: req.Interval,
	}
	err = models.CreateWeixinKeywordSetting(setting)
	if err != nil {
		holmes.Error("create weixin keyword setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) CreateWeixinKeyword(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinKeyword{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinKeyword json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateWeixinKeyword(req)
	if err != nil {
		holmes.Error("create weixin keyword error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetWeixinSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.Weixin{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetWeixinSetting json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	holmes.Debug("get weixin setting req: %v", req)

	weixin := &models.Weixin{
		Wechat: req.Wechat,
	}
	has, err := models.GetWeixin(weixin)
	if err != nil {
		holmes.Error("get weixin error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if !has {
		holmes.Error("Has no this weixin[%s]", req.Wechat)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if weixin.WxId != req.WxId || weixin.NickName != req.NickName {
		weixin.WxId = req.WxId
		weixin.NickName = req.NickName
		err = models.UpdateWeixinWxid(weixin)
		if err != nil {
			holmes.Error("update weixin error: %v", err)
		}
	}

	setting := &proto.WeixinSetting{
		WeixinId: weixin.ID,
	}

	verifySetting, err := models.GetWeixinVerifySettingDetail(weixin.ID)
	if err != nil {
		holmes.Error("get weixin verify setting error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if verifySetting != nil {
		setting.Verify.IfAutoVerified = verifySetting.IfAutoVerified
		setting.Verify.Interval = verifySetting.Interval
		err = json.Unmarshal([]byte(verifySetting.Reply), &setting.Verify.Reply)
		if err != nil {
			holmes.Error("json unmarshal verify setting reply error: %v", err)
		}
	}

	keywordSettingList, err := models.GetWeixinKeywordSettingList(weixin.ID)
	if err != nil {
		holmes.Error("get weixin keyword setting list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	for _, v := range keywordSettingList {
		kSetting := proto.KeywordSetting{
			ChatType: v.ChatType,
			MsgType:  v.MsgType,
			Keyword:  v.Keyword,
			Interval: v.Interval,
		}
		err = json.Unmarshal([]byte(v.Reply), &kSetting.Reply)
		if err != nil {
			holmes.Error("json unmarshal keyword setting reply error: %v", err)
		} else {
			setting.Keyword = append(setting.Keyword, kSetting)
		}
	}

	rsp.Data = setting
}

func (self *Logic) GetAllWeixin(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	list, err := models.GetAllWeixinList()
	if err != nil {
		holmes.Error("get all weixin error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) GetAllVerifySetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	list, err := models.GetAllVerifyList()
	if err != nil {
		holmes.Error("get all weixin verify list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) GetAllKeywordSetting(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	list, err := models.GetAllKeywordList()
	if err != nil {
		holmes.Error("get all weixin keyword list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}
