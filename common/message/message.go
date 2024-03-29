package message

const (
	LoginMsgType  = "LoginMsgType"
	LoginResMsgType  = "LoginResMsgType"
	LogoutMsgType  = "LogoutMsgType"
	LogoutResMsgType = "LogoutResMsgType"
	RegisterMsgType  = "RegisterMsgType"
	RegisterResMsgType  = "RegisterResMsgType"
	NotifyUserStatusMsgType = "NotifyUserStatusMsgType"
	SmsMsgType = "SmsMsgType"
	PrivateChatSmsMsgType = "PrivateChatSmsMsgType"
)

const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

type LoginResMsg struct {
	Code int `json:"code"`
	UsersId []int  `json:"users_id"`
	ErrorInfo string `json:"error_info"`
}

type LogoutMsg struct {
	UserId int `json:"user_id"`
}

type LogoutResMsg struct {
	Code int `json:"code"`
	ErrorInfo string `json:"error_info"`
}

type RegisterMsg struct {
	User User `json:"user"`
	//UserId int `json:"user_id"`
	//UserPwd string `json:"user_pwd"`
	//UserName string `json:"user_name"`
}

type RegisterResMsg struct {
	Code int `json:"code"`
	ErrorInfo string `json:"error_info"`
}

type NotifyUserStatusMsg struct {
	UserId int `json:"user_id"`
	UserStatus int `json:"user_status"`
}

type SmsMsg struct {
	User
	Content string `json:"content"`
}

type PrivateChatSmsMsg struct {
	SmsMsg
	DestUserId int `json:"dest_user_id"`
}