package proto

type ReqID struct {
	Id int64 `json:"id"`
}

type Weixin struct {
	WxId       string  `json:"wxId"`
	Wechat     string  `json:"wechat"`
	NickName   string  `json:"nickName"`
	WeixinId   int64   `json:"weixinId"`
	VerifyId   int64   `json:"verifyId"`
	KeywordIds []int64 `json:"keywordIds"`
	TaskIds    []int64 `json:"taskIds"`
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
	WeixinId int64            `json:"weixinId"`
	Verify   VerifySetting    `json:"verifySetting"`
	Keyword  []KeywordSetting `json:"keywordSetting"`
}

// task
type BatchTaskList struct {
	TaskIds []int64 `json:"taskIds"`
	Weixins []int64 `json:"weixins"`
}

type Task struct {
	TaskType  int64       `json:"taskId"`
	IfDefault int64       `json:"IfDefault"`
	Data      interface{} `json:"data"`
}

type LinkMsg struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	LinkUrl    string `json:"linkUrl"`
	ShowPicUrl string `json:"showPicUrl"`
}

type ContactsMass struct {
	Interval     int64     `json:"interval"`
	ControlLimit int64     `json:"controlLimit"`
	TextMsgs     []string  `json:"textMsgs"`
	CardMsgs     []string  `json:"cardMsgs"`
	PicMsg       string    `json:"picMsg"`
	LinkMsgs     []LinkMsg `json:"linkMsgs"`
}

type FriendsCircle struct {
	Text  string   `json:"text"`
	Type  string   `json:"type"`
	Media []string `json:"media"`
}

type AttentionCard struct {
	Cards    []string `json:"cards"`
	Interval int64    `json:"interval"`
}

type HeadImg struct {
	HeadUrl       string `json:"headUrl"`
	BackgroundUrl string `json:"backgroundUrl"`
}

type BasicInfo struct {
	Sex       int    `json:"sex"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Signature string `json:"signature"`
	Nickname  string `json:"nickname"`
}

type WxUserInfo struct {
	HI HeadImg   `json:"headImg"`
	BI BasicInfo `json:"basicInfo"`
}

// contact
type Contact struct {
	UserName    string `json:"userName"`
	AliasName   string `json:"aliasName"`
	NickName    string `json:"nickName"`
	PhoneNumber string `json:"phoneNumber"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Sex         int64  `json:"sex"`
	Remark      string `json:"remark"`
}

type SyncContacts struct {
	Myself      string    `json:"myself"`
	ContactData []Contact `json:"contactData"`
}

type AddContact struct {
	Myself      string  `json:"myself"`
	ContactData Contact `json:"contactData"`
}
