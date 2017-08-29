package ext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/reechou/holmes"
	"github.com/reechou/weixin-x/config"
)

type AlarmExt struct {
	cfg    *config.Config
	client *http.Client
}

func NewAlarmExt(cfg *config.Config) *AlarmExt {
	return &AlarmExt{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (self *AlarmExt) DoAlarm(request *AlarmReq) error {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		holmes.Error("json encode error: %v", err)
		return err
	}
	
	url := "http://" + self.cfg.AlarmBackend.Host + ALARM_URI_DO
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		holmes.Error("http new request error: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := self.client.Do(req)
	if err != nil {
		holmes.Error("http do request error: %v", err)
		return err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		holmes.Error("ioutil ReadAll error: %v", err)
		return err
	}
	var response AlarmRsp
	err = json.Unmarshal(rspBody, &response)
	if err != nil {
		holmes.Error("json decode error: %v [%s]", err, string(rspBody))
		return err
	}
	if response.Code != 0 {
		holmes.Error("alarm error: code[%d]", response.Code)
		return fmt.Errorf("alarm error.")
	}
	
	return nil
}
