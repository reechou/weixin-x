package ext

const (
	ALARM_URI_DO = "/api/doalarm"
)

const (
	ALARM_CHANNEL_SMS = "sms"
	ALARM_CHANNEL_GZH = "gzh"
)

const (
	ALARM_TYPE_WARN = iota
	ALARM_TYPE_SYS_ERROR
)

type AlarmRsp struct {
	Code int64 `json:"code"`
}

type AuthorizationReq struct {
	OpenId   string `json:"openid"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AlarmReq struct {
	AlarmType    int64    `json:"alarmType"`
	AlarmTime    int64    `json:"alarmTime"`
	AlarmMsg     string   `json:"alarmMsg"`
	AlarmRemark  string   `json:"alarmRemark"`
	AlarmHost    string   `json:"alarmHost"`
	AlarmService string   `json:"alarmService"`
	AlarmStatus  string   `json:"alarmStatus"`
	Channels     []string `json:"channels"`
	ToUsers      []string `json:"toUsers"`
}
