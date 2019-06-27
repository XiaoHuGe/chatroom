package message

const (
	LoginMsgType  = "LoginMsgType"
	LoginResMsgType  = "LoginResMsgType"
	RegisterMsgType  = "RegisterMsgType"
	RegisterResMsgType  = "RegisterResMsgType"
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