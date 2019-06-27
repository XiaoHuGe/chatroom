package message

type User struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
	UserStatus int `json:"User_status"`
	Sex string `json:"sex"`
}
