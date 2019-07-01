package services

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"chatroom/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	userId int
	UserStatus int
}

func marshalMessage(userId, userStatus int) (data []byte, err error) {
	// 创建Message结构体对象
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType
	// 创建NotifyOnlineMsg结构体对象
	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.UserStatus = userStatus

	// 把NotifyOnlineMsg序列化
	data, err = json.Marshal(notifyUserStatusMsg)
	if err != nil {
		return
	}
	msg.Data = string(data)

	// 把Message序列化
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	return
}

// 通知其它在线用户我状态发生改变
func (this *UserProcess)NotifyOtherOnlineUser(userId, userStatus int) {

	data, err := marshalMessage(userId, userStatus)
	if err != nil {
		fmt.Println("json.Marsha loginMsg error: ", err)
		return
	}
	// 遍历所有在线用户
	for id, up := range userMgr.users{
		if id == userId {
			continue
		}
		// 离线用户
		if up.UserStatus == message.UserOffline {
			continue
		}
		// 注意：使用up调用
		//up.NotifyMeOnline(data)
		// 发送data到客户端
		transfer := &utils.Transfer{
			Conn: up.Conn,
		}
		err := transfer.WritePkg(data)
		if err != nil {
			fmt.Println("utils.WritePkg error:", err)
			return
		}
	}
}

func (this *UserProcess)ServerProcessRegister(msg *message.Message) (err error){

	// 创建RegisterMsg结构体
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	fmt.Println("LoginMsg userID:", registerMsg.User.UserId)
	fmt.Println("LoginMsg userPwd:", registerMsg.User.UserPwd)
	fmt.Println("LoginMsg userPwd:", registerMsg.User.UserName)

	// 创建Message结构体
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType

	// 创建RegisterResMsg结构体
	var registerResMsg message.RegisterResMsg

	// 判断是否注册成功
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_UESR_EXISTS {
			fmt.Println("用户已存在")
			registerResMsg.Code = 501
			registerResMsg.ErrorInfo = err.Error()
		} else {
			fmt.Println("服务器内部错误")
			registerResMsg.Code = 505
			registerResMsg.ErrorInfo = "服务器内部错误"
		}
	} else {
		// 把用户添加到map
		this.userId = registerMsg.User.UserId
		this.UserStatus = message.UserOffline
		userMgr.AddOnlineUser(this)

		registerResMsg.Code = 200
		fmt.Println("注册成功")
	}

	// 序列化loginResMsg
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("json.Marshal error: ", err)
	}
	//
	resMsg.Data = string(data)

	// 序列化Message
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marsha Message error: ", err)
	}
	transfer := &utils.Transfer{
		Conn:this.Conn,
	}
	transfer.WritePkg(data)
	return
}


func (this *UserProcess)ServerProcessLogin(msg *message.Message) (err error){

	// 创建LoginMsg结构体
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	fmt.Println("LoginMsg userID:", loginMsg.UserId)
	fmt.Println("LoginMsg userPwd:", loginMsg.UserPwd)

	// 创建Message结构体
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	// 创建LoginResMsg结构体
	var loginResMsg message.LoginResMsg

	// 判断登录
	_, err = model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_UESR_NOTEXISTS {
			fmt.Println("用户不存在")
			loginResMsg.Code = 500
			loginResMsg.ErrorInfo = err.Error()
		} else if err == model.ERROR_UESR_PWD {
			fmt.Println("用户密码错误")
			loginResMsg.Code = 403
			loginResMsg.ErrorInfo = err.Error()
		} else {
			fmt.Println("服务器内部错误")
			loginResMsg.Code = 505
			loginResMsg.ErrorInfo = "服务器内部错误"
		}
	} else {
		loginResMsg.Code = 200
		fmt.Println("登录成功")

		// 更新map用户状态
		this.userId = loginMsg.UserId
		this.UserStatus = message.UserOnline
		userMgr.UpdateOnlineUser(this)

		// 返回给客户端所有在线用户
		for k, up := range userMgr.users{
			if up.UserStatus == message.UserOnline {
				loginResMsg.UsersId = append(loginResMsg.UsersId, k)
			}
			//loginResMsg.UsersId = append(loginResMsg.UsersId, k)
			//println("在线用户：", k)
		}

		// 实时通知其他用户上线消息
		this.NotifyOtherOnlineUser(loginMsg.UserId, message.UserOnline)
	}

	// 序列化loginResMsg
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal error: ", err)
	}
	resMsg.Data = string(data)

	// 序列化Message
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marsha Message error: ", err)
	}
	transfer := &utils.Transfer{
		Conn:this.Conn,
	}
	transfer.WritePkg(data)
	return
}

func (this *UserProcess)ServerProcessLogout(msg *message.Message) (err error){

	// 创建LogoutMsg结构体
	var logoutMsg message.LogoutMsg
	err = json.Unmarshal([]byte(msg.Data), &logoutMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	fmt.Println("LogoutMsg userID:", logoutMsg.UserId)

	// 创建Message结构体
	var resMsg message.Message
	resMsg.Type = message.LogoutResMsgType

	// 创建LogoutResMsg结构体
	var logoutResMsg message.LogoutResMsg

	// 判断登录
	_, err = model.MyUserDao.Logout(logoutMsg.UserId)
	if err != nil {
		if err == model.ERROR_UESR_NOTEXISTS {
			fmt.Println("用户不存在")
			logoutResMsg.Code = 500
			logoutResMsg.ErrorInfo = err.Error()
		} else {
			fmt.Println("服务器内部错误")
			logoutResMsg.Code = 505
			logoutResMsg.ErrorInfo = "服务器内部错误"
		}
	} else {
		logoutResMsg.Code = 200
		fmt.Println("退出成功")

		// 更新用户状态
		this.userId = logoutMsg.UserId
		this.UserStatus = message.UserOffline
		userMgr.UpdateOnlineUser(this)
		//userMgr.DeleteOnlineUser(this)

		// 实时通知其他用户退出消息
		this.NotifyOtherOnlineUser(logoutMsg.UserId, message.UserOffline)
	}

	// 序列化loginResMsg
	data, err := json.Marshal(logoutResMsg)
	if err != nil {
		fmt.Println("json.Marshal error: ", err)
	}
	resMsg.Data = string(data)

	// 序列化Message
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marsha Message error: ", err)
	}
	transfer := &utils.Transfer{
		Conn:this.Conn,
	}
	transfer.WritePkg(data)
	return
}