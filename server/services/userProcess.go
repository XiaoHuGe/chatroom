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
	transfer := utils.Transfer{
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
	}

	// 序列化loginResMsg
	data, err := json.Marshal(loginResMsg)
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
	transfer := utils.Transfer{
		Conn:this.Conn,
	}
	transfer.WritePkg(data)
	return
}
