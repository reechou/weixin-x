package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/proto"
)

func (self *Logic) CreateLiebianType(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.LiebianType{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateLiebianType json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateLiebianType(req)
	if err != nil {
		holmes.Error("create liebian type error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetLiebianTypeList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	list, err := models.GetLiebianTypeList()
	if err != nil {
		holmes.Error("get liebian type list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) DeleteLiebianType(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.LiebianType{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteLiebianType json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.DelLiebianType(req)
	if err != nil {
		holmes.Error("delete liebian type error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) CreateWeixinGroup(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinGroup{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinGroup json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateWeixinGroup(req)
	if err != nil {
		holmes.Error("create weixin group error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) DeleteWeixinGroup(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinGroup{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteWeixinGroup json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.DelWeixinGroup(req)
	if err != nil {
		holmes.Error("delete weixin group error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetWeixinGroupList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	list, err := models.GetWeixinGroupList()
	if err != nil {
		holmes.Error("get weixin group list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) CreateWeixinGroupMemberList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	var req []models.WeixinGroupMember
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		holmes.Error("CreateWeixinGroupMemberList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateWeixinGroupMemberList(req)
	if err != nil {
		holmes.Error("create weixin group member list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) DeleteWeixinGroupMember(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinGroupMember{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteWeixinGroupMember json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.DelWeixinGroupMember(req)
	if err != nil {
		holmes.Error("delete weixin group member error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetWeixinGroupMemberList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.GetWeixinGroupMemberListReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetWeixinGroupMemberList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	list, err := models.GetWeixinGroupMemberDetailList(req.GroupId)
	if err != nil {
		holmes.Error("delete weixin group member error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) CreateLiebianPool(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	var req []models.LiebianPool
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		holmes.Error("CreateLiebianPool json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateLiebianPoolList(req)
	if err != nil {
		holmes.Error("create liebian pool list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) DeleteLiebianPool(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	var req []int64
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		holmes.Error("DeleteLiebianPool json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.DelLiebianPoolList(req)
	if err != nil {
		holmes.Error("delete liebian pool list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetLiebianPool(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.GetLiebianPoolReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetLiebianPool json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	list, err := models.GetLiebianPoolWeixinList(req.LiebianType)
	if err != nil {
		holmes.Error("delete liebian pool list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) GetUserLiebianInfo(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.GetLiebianInfoReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetUserLiebianInfo json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	qrcodeBind := &models.QrcodeBind{
		AppId:       req.AppId,
		OpenId:      req.OpenId,
		LiebianType: req.LiebianType,
	}
	has, err := models.GetQrcodeBind(qrcodeBind)
	if err != nil {
		holmes.Error("get qrcode bind error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	result := new(proto.GetLiebianInfoRsp)
	if has {
		if qrcodeBind.BindQrcode != "" {
			result.Qrcode = qrcodeBind.BindQrcode
			return
		}
	}

	liebianList, err := models.GetLiebianPoolWeixinList(req.LiebianType)
	if err != nil {
		holmes.Error("get liebian pool weixin list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if len(liebianList) == 0 {
		holmes.Error("get liebian pool weixin list of type[%d] is nil", req.LiebianType)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("get qrcode list from type is nil")
		return
	}
	offset := rand.Intn(len(liebianList))
	result.Qrcode = liebianList[offset].Weixin.QrcodeUrl
	if result.Qrcode == "" {
		holmes.Error("get liebian pool of weixin[%v] qrcodeurl is nil", liebianList[offset])
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("get qrcode of url offset is nil")
		return
	}

	qrcodeBind.BindQrcode = result.Qrcode
	err = models.CreateQrcodeBind(qrcodeBind)
	if err != nil {
		holmes.Error("create qrcode bind error: %v", err)
	}

}

func init() {
	rand.Seed(time.Now().UnixNano())
}
