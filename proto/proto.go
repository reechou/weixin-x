package proto

type Weixin struct {
	WxId     string `json:"wxId"`
	Wechat   string `json:"wechat"`
	NickName string `json:"nickName"`
}

type MsgInfo struct {
	MsgType   int    `json:"msgType"`
	Content   string `json:"content"`
	MsgSource string `json:"msgSource"`
}

type VerifySetting struct {
	IfAutoVerified bool      `json:"ifAutoVerified"`
	Reply          []MsgInfo `json:"reply"`
	Interval       int64     `json:"interval"`
}

type KeywordSetting struct {
	ChatType string    `json:"chatType"`
	MsgType  int64     `json:"msgType"`
	Keyword  string    `json:"keyword"`
	Reply    []MsgInfo `json:"reply"`
	Interval int64     `json:"interval"`
}

type WeixinSetting struct {
	Verify  VerifySetting    `json:"verifySetting"`
	Keyword []KeywordSetting `json:"keywordSetting"`
}
