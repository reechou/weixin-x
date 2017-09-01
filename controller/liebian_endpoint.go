package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"sort"

	"github.com/jinzhu/now"
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

func (self *Logic) UpdateLiebianTypeLimit(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.LiebianType{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("UpdateLiebianTypeLimit json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.UpdateLiebianTypeLimit(req)
	if err != nil {
		holmes.Error("update liebian type limit error: %v", err)
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
	
	if len(req) == 0 {
		return
	}

	err := models.CreateLiebianPoolList(req)
	if err != nil {
		holmes.Error("create liebian pool list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	var ids []int64
	for _, v := range req {
		ids = append(ids, v.WeixinId)
	}
	msgStr := fmt.Sprintf("裂变TYPE[%d] 上线资源池成员[%v]", req[0].LiebianType, ids)
	oprMsg := &models.LiebianOprMsg{
		LiebianType: req[0].LiebianType,
		Msg:         msgStr,
	}
	models.CreateLiebianOprMsg(oprMsg)
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
	holmes.Debug("delete liebian pool ids: %v", req)
	
	liebianPoolList, err := models.GetLiebianPoolListFromIds(req)
	if err != nil {
		holmes.Error("get liebian pool list from ids error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err = models.DelLiebianPoolList(req)
	if err != nil {
		holmes.Error("delete liebian pool list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	// create opr msg
	if len(liebianPoolList) == 0 {
		return
	}
	var ids []int64
	for _, v := range liebianPoolList {
		ids = append(ids, v.WeixinId)
	}
	msgStr := fmt.Sprintf("裂变TYPE[%d] 下线资源池成员[%v]", liebianPoolList[0].LiebianType, ids)
	oprMsg := &models.LiebianOprMsg{
		LiebianType: liebianPoolList[0].LiebianType,
		Msg:         msgStr,
	}
	models.CreateLiebianOprMsg(oprMsg)
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
	todayZero := now.BeginningOfDay().Unix()
	for i := 0; i < len(list); i++ {
		if list[i].Weixin.LastAddContactTime < todayZero {
			list[i].Weixin.TodayAddContactNum = 0
		}
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
	var listOffset int
	start := time.Now()
	defer func() {
		holmes.Debug("get user liebian req[%v] rsp[offset-%d][%v] end, use time: %v.", req, listOffset, rsp.Data, time.Now().Sub(start))
	}()
	
	// collect pv
	self.dsw.Collect(&StatisticsDataInfo{TypeId: int64(models.S_DATA_LIEBIAN_PV), Data: 1, LiebianType: req.LiebianType})

	qrcodeBind := &models.QrcodeBind{
		AppId:       req.AppId,
		OpenId:      req.OpenId,
		LiebianType: req.LiebianType,
	}
	hasBind, err := models.GetQrcodeBind(qrcodeBind)
	if err != nil {
		holmes.Error("get qrcode bind error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	result := new(proto.GetLiebianInfoRsp)
	if hasBind {
		if qrcodeBind.BindQrcode != "" {
			result.Qrcode = qrcodeBind.BindQrcode
			rsp.Data = result
			return
		}
	}
	
	// collect uv
	self.dsw.Collect(&StatisticsDataInfo{TypeId: int64(models.S_DATA_LIEBIAN_UV), Data: 1, LiebianType: req.LiebianType})

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
	listOffset = rand.Intn(len(liebianList))
	result.Qrcode = liebianList[listOffset].Weixin.QrcodeUrl
	if result.Qrcode == "" {
		holmes.Error("get liebian pool of weixin[%v] qrcodeurl is nil", liebianList[listOffset])
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("get qrcode of url offset is nil")
		return
	}

	qrcodeBind.BindQrcode = result.Qrcode
	qrcodeBind.WeixinId = liebianList[listOffset].Weixin.ID
	err = models.CreateQrcodeBind(qrcodeBind)
	if err != nil {
		holmes.Error("create qrcode bind error: %v", err)
	}

	rsp.Data = result
}

func (self *Logic) GetDataStatistical(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.GetDataStatisticalReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetDataStatistical json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	list, err := models.GetStatisticalDataList(req.TypeId, req.LiebianType, req.StartTime, req.EndTime)
	if err != nil {
		holmes.Error("get statistical data error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if req.LiebianType == 0 {
		dataMap := make(map[int]*models.StatisticalData)
		for _, v := range list {
			dataV, ok := dataMap[int(v.TimeSeries)]
			if ok {
				dataV.Data = dataV.Data + v.Data
			} else {
				dataMap[int(v.TimeSeries)] = &models.StatisticalData{
					TypeId:     v.TypeId,
					Data:       v.Data,
					TimeSeries: v.TimeSeries,
				}
			}
		}
		var keys []int
		for k := range dataMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var result []*models.StatisticalData
		for _, k := range keys {
			result = append(result, dataMap[k])
		}
		rsp.Data = result

		//var mergeList []models.StatisticalData
		//var nowTS int64 = -1
		//idx := -1
		//for _, v := range list {
		//	if nowTS != v.TimeSeries {
		//		idx++
		//		nowTS = v.TimeSeries
		//		mergeList = append(mergeList, models.StatisticalData{
		//			TypeId:     v.TypeId,
		//			Data:       v.Data,
		//			TimeSeries: v.TimeSeries,
		//		})
		//	} else {
		//		mergeList[idx].Data = mergeList[idx].Data+v.Data
		//	}
		//}
		//rsp.Data = mergeList
		return
	}
	rsp.Data = list
}

func (self *Logic) GetLiebianErrorMsgList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.GetLiebianErrorMsgReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetLiebianErrorMsgList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	list, err := models.GetLiebianErrorMsgList(req.LiebianType)
	if err != nil {
		holmes.Error("get liebian error msg list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func (self *Logic) GetLiebianOprMsgList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.GetLiebianErrorMsgReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetLiebianOprMsgList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	list, err := models.GetLiebianOprMsgList(req.LiebianType)
	if err != nil {
		holmes.Error("get liebian opr msg list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
