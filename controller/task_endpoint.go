package controller

import (
	"encoding/json"
	"net/http"

	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/models"
	"github.com/reechou/weixin-x/proto"
)

func (self *Logic) CreateTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.Task{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateTask json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	data, err := json.Marshal(req.Data)
	if err != nil {
		holmes.Error("json marshal data error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	task := &models.WeixinTask{
		TaskType:  req.TaskType,
		IfDefault: req.IfDefault,
		Data:      string(data),
	}
	err = models.CreateWeixinTask(task)
	if err != nil {
		holmes.Error("create weixin task error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) BatchCreateTaskList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.BatchTaskList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("BatchCreateTaskList json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	var list []models.WeixinTaskList
	for _, v := range req.TaskIds {
		for _, v2 := range req.Weixins {
			list = append(list, models.WeixinTaskList{
				WeixinId:     v2,
				WeixinTaskId: v,
			})
		}
	}
	if list != nil || len(list) != 0 {
		err := models.CreateWeixinTaskInfoList(list)
		if err != nil {
			holmes.Error("create weixin task error: %v", err)
			rsp.Code = proto.RESPONSE_ERR
			return
		}
	}
}

func (self *Logic) DeleteTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.ReqID{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("DeleteWeixinTask json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.DelWeixinTask(&models.WeixinTask{ID: req.Id})
	if err != nil {
		holmes.Error("delete weixin task error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) UpdateTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &models.WeixinTask{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("UpdateTask json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	err := models.UpdateWeixinTask(req)
	if err != nil {
		holmes.Error("update weixin task error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) GetTaskFromId(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.ReqID{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetTaskFromId json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	
	task := &models.WeixinTask{
		ID: req.Id,
	}
	has, err := models.GetWeixinTaskFromId(task)
	if err != nil {
		holmes.Error("get weixin task error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	if !has {
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = "can not found"
		return
	}
	rsp.Data = task
}

func (self *Logic) CreateWeixinTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &models.WeixinTaskList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateWeixinTask json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	err := models.CreateWeixinTaskList(req)
	if err != nil {
		holmes.Error("create weixin task list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
}

func (self *Logic) transferTask(taskType int64, data string) interface{} {
	switch taskType {
	case proto.TASK_ID_CONTACTS_MASS:
		task := &proto.ContactsMass{}
		err := json.Unmarshal([]byte(data), task)
		if err != nil {
			holmes.Error("json unmarshal error: %v", err)
			return nil
		}
		return task
	case proto.TASK_ID_FRIENDS_CIRCLE:
		task := &proto.FriendsCircle{}
		err := json.Unmarshal([]byte(data), task)
		if err != nil {
			holmes.Error("json unmarshal error: %v", err)
			return nil
		}
		return task
	case proto.TASK_ID_ATTENTION_CARD:
		task := &proto.AttentionCard{}
		err := json.Unmarshal([]byte(data), task)
		if err != nil {
			holmes.Error("json unmarshal error: %v", err)
			return nil
		}
		return task
	case proto.TASK_ID_MODIFY_USERINFO:
		task := &proto.WxUserInfo{}
		err := json.Unmarshal([]byte(data), task)
		if err != nil {
			holmes.Error("json unmarshal error: %v", err)
			return nil
		}
		return task
	}

	return nil
}

func (self *Logic) GetTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.Weixin{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetTask json decode error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	//holmes.Debug("get task req: %v", req)

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

	var taskList []proto.Task
	if weixin.IfExecDefaultTask == 0 {
		wxTask, err := models.GetWeixinDefaultTaskList()
		if err != nil {
			holmes.Error("get weixin default task list error: %v", err)
			rsp.Code = proto.RESPONSE_ERR
			return
		}
		for _, v := range wxTask {
			holmes.Debug("default task: %v", v)
			task := self.transferTask(v.TaskType, v.Data)
			if task != nil {
				taskList = append(taskList, proto.Task{
					TaskType:  v.TaskType,
					IfDefault: v.IfDefault,
					Data:      task,
				})
			}
		}
		// update if default task
		weixin.IfExecDefaultTask = 1
		err = models.UpdateWeixinIfExecDefaultTask(weixin)
		if err != nil {
			holmes.Error("update weixin if exec default task error: %v", err)
		}
	}

	wxTask, err := models.GetWeixinTaskList(weixin.ID)
	if err != nil {
		holmes.Error("get weixin task list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}

	if wxTask == nil || len(wxTask) == 0 {
		rsp.Data = taskList
		return
	}

	for _, v := range wxTask {
		task := self.transferTask(v.TaskType, v.Data)
		if task != nil {
			taskList = append(taskList, proto.Task{
				TaskType:  v.TaskType,
				IfDefault: v.IfDefault,
				Data:      task,
			})
		}
	}

	err = models.UpdateWeixinTaskListFromWeixinId(weixin.ID)
	if err != nil {
		holmes.Error("update weixin if exec task error: %v", err)
	}
	rsp.Data = taskList
}

func (self *Logic) GetAllTask(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	list, err := models.GetAllTaskList()
	if err != nil {
		holmes.Error("get all weixin task list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		return
	}
	rsp.Data = list
}
