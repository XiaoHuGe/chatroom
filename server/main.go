package main

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func serverProcessLogin(conn net.Conn, msg *message.Message) (err error){

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
	if loginMsg.UserId == 100 &&  loginMsg.UserPwd == "123456"{
		loginResMsg.Code = 200
	} else {
		loginResMsg.Code = 500
		loginResMsg.ErrorInfo = "用户名或密码不正确"
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
	utils.WritePkg(conn, data)
	return
}

func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginMsgType:
			 err = serverProcessLogin(conn, msg)
		case message.RegisterMsgType:
		//
		default:
			println("消息类型不存在")
	}
	return
}
func process(conn net.Conn) {

	defer conn.Close()

	for {
		// 读取客户端消息
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出....")
				return
			} else {
				fmt.Println("readPkg error: ", err)
				return
			}
		}
		fmt.Println("msg :", msg)

		// 处理客户端消息
		serverProcessMsg(conn, &msg)
	}
}

func main() {

	fmt.Println("服务器在8899端口监听...")
	listener, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		fmt.Println("net.Listen error: ", err)
	}
	defer listener.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error: ", err)
		}
		// 连接成功，开启一个协程保持与客户端的连接
		process(conn)
	}
}