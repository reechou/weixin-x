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
	BindId         int64     `json:"bindId"`
	IfAutoVerified bool      `json:"ifAutoVerified"`
	Reply          []MsgInfo `json:"reply"`
	Interval       int64     `json:"interval"`
}

type KeywordSetting struct {
	BindId   int64     `json:"bindId"`
	ChatType string    `json:"chatType"`
	MsgType  int64     `json:"msgType"`
	Keyword  string    `json:"keyword"`
	Reply    []MsgInfo `json:"reply"`
	Interval int64     `json:"interval"`
}

type WeixinSetting struct {
	WeixinId                 int64            `json:"weixinId"`
	Verify                   VerifySetting    `json:"verifySetting"`
	Keyword                  []KeywordSetting `json:"keywordSetting"`
	RestartWithUnreceivedMsg int64            `json:"restartWithUnreceivedMsg"`
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

type ScanQrcode struct {
	Interval   int64    `json:"interval"`
	QrcodeUrls []string `json:"qrcodeUrls"`
}

type LinkMsg struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	LinkUrl    string `json:"linkUrl"`
	ShowPicUrl string `json:"showPicUrl"`
	DataUrl    string `json:"dataUrl"`
}

type ContactsMass struct {
	Interval     int64     `json:"interval"`
	ControlLimit int64     `json:"controlLimit"`
	TextMsgs     []string  `json:"textMsgs"`
	CardMsgs     []string  `json:"cardMsgs"`
	PicMsg       string    `json:"picMsg"`
	LinkMsgs     []LinkMsg `json:"linkMsgs"`
	Friends      []string  `json:"friends"`
}

type FriendsCircle struct {
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Media   []string `json:"media"`
	Title   string   `json:"title"`   // for link wc
	LinkUrl string   `json:"linkUrl"` // for link wc
	PicUrl  string   `json:"picUrl"`  // for link wc
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

type WeixinContactBindReq struct {
	Myself string `json:"myself"`
	WxId   string `json:"wxId"`
	CardId string `json:"cardId"`
}

type WeixinContactBindRsp struct {
	BindCard interface{} `json:"bindCard"`
}

type GetFriendsReq struct {
	WeixinId int64 `json:"weixinId"`
	Offset   int64 `json:"offset"`
	Num      int64 `json:"num"`
}

type GetFriendsFromTimeReq struct {
	WeixinId  int64 `json:"weixinId"`
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type GetFriendsFromTagReq struct {
	WeixinId  int64 `json:"weixinId"`
	TagId     int64 `json:"tagId"`
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type CreateSelectedFriendsTaskReq struct {
	WeixinId     int64    `json:"weixinId"`
	WeixinTaskId int64    `json:"weixinTaskId"`
	TagId        int64    `json:"tagId"`
	Friends      []string `json:"friends"`
}

type GetTimerTaskListFromWeixinReq struct {
	WeixinId int64 `json:"weixinId"`
}

type GetWeixinGroupMemberListReq struct {
	GroupId int64 `json:"groupId"`
}

// liebian
type GetLiebianPoolReq struct {
	LiebianType int64 `json:"liebianType"`
}

type GetLiebianInfoReq struct {
	LiebianType int64  `json:"liebianType"`
	AppId       string `json:"appId"`
	OpenId      string `json:"openId"`
}

type GetLiebianInfoRsp struct {
	Status int64  `json:"status"`
	Qrcode string `json:"qrcode"`
}

type GetResourcePoolReq struct {
	WxType int64 `json:"wxType"`
}

// data statistic
type GetDataStatisticalReq struct {
	TypeId      int64 `json:"typeId"`
	LiebianType int64 `json:"liebianType"`
	StartTime   int64 `json:"startTime"`
	EndTime     int64 `json:"endTime"`
}

type GetLiebianErrorMsgReq struct {
	LiebianType int64 `json:"liebianType"`
}

// chat room
type ChatroomFullSetting struct {
	ChatroomSetTimeAfterFull int64  `json:"chatroomSetTimeAfterFull"`
	ChatroomFullMemberNum    int64  `json:"chatroomFullMemberNum"`
	ChatroomNotice           string `json:"chatroomNotice"`
	IfAutoChangeAccess       int64  `json:"ifAutoChangeAccess"`
	ChatroomChangeOwner      string `json:"chatroomChangeOwner"`
	IfAutoShowInContactBook  int64  `json:"ifAutoShowInContactBook"`
}

type ChatroomCommonSetting struct {
	WelcomeReply            []MsgInfo           `json:"welcomeReply"`    // 新人进群
	ScreenshotReply         []MsgInfo           `json:"screenshotReply"` // 截图
	SysLoopReply            []MsgInfo           `json:"sysLoopReply"`    // 循环
	Interval                int64               `json:"interval"`
	WelcomeReplyLoopTime    int64               `json:"welcomeReplyLoopTime"`
	ScreenshotReplyLoopTime int64               `json:"screenshotReplyLoopTime"`
	SysLoopTime             int64               `json:"sysLoopTime"`
	ChatroomRole            string              `json:"chatroomRole"` // chatroom-role-master: 主 chatroom-role-slave: 备
	FullSetting             ChatroomFullSetting `json:"chatroomFullSetting"`
}
