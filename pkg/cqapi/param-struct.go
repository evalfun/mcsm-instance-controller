package cqapi

type Report interface{}

type CommonReport struct {
	Time   int64 `json:"time"`
	SelfID int64 `json:"self_id"`
	// message request notice meta_event
	PostType string `json:"post_type"`
}

type MessageReportSender struct {
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserID   int64  `json:"user_id"`
}

// 消息上报
type MessageReport struct {
	CommonReport
	//private, group
	MessageType string `json:"message_type"`
	//group, public normal
	SubType    string              `json:"sub_type"`
	MessageID  int32               `json:"message_id"`
	UserID     int64               `json:"user_id"`
	Font       int                 `json:"font"`
	Sender     MessageReportSender `json:"sender"`
	Anonymous  string              `json:"anonymous"`
	MessageSeq int64               `json:"message_seq"`
	Message    string              `json:"message"`
	RawMessage string              `json:"raw_message"`
	GroupID    int64               `json:"group_id"`
}

// 请求上报
type RequestReport struct {
	CommonReport
	RequestType string `json:"request_type"`
}

// 通知上报
type NoticeReport struct {
	CommonReport
	NoticeType string `json:"notice_type"`
}

// 元事件上报
type MetaEventReport struct {
	CommonReport
	MetaEventType string `json:"meta_event_type"`
}

type WebsocketJsonRequestParams struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

type SendPrivateMessageParams struct {
	UserID     int64  `json:"user_id"`
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type SendGroupMessageParams struct {
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type SendMessageParams struct {
	//private group
	MessageType string `json:"message_type"`
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  bool   `json:"auto_escape"`
}
