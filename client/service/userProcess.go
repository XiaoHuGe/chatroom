package service

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess)Register(userId int, userPwd string, userName string) (err error) {
	fmt.Println("进入注册系统...")
	fmt.Println("用户ID：", userId)
	fmt.Println("用户密码：", userPwd)
	fmt.Println("用户昵称：", userName)

	// 1.连接到服务器
	conn , err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		fmt.Println("net.Dial error")
	}
	defer conn.Close()

	// 2. 通过conn发送消息
	// 创建Message结构体对象
	var msg message.Message
	msg.Type = message.RegisterMsgType
	// 创建LoginMsg结构体对象
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	// 把LoginMsg序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marsha loginMsg error")
		return
	}
	msg.Data = string(data)

	// 把Message序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marsha Message error")
		return
	}

	// 发送data到服务器
	transfer := &utils.Transfer{
		Conn:conn,
	}
	err = transfer.WritePkg(data)
	//err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("utils.WritePkg error:", err)
		return
	}

	// 3. 获取服务器返回内容
	//msg, err = utils.ReadPkg(conn)
	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) error")
		return
	}

	var registerResMsg message.RegisterResMsg

	// 反序列化message
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	if registerResMsg.Code  == 200 {
		fmt.Println("注册成功... 请重新登录")
		fmt.Println("code:", registerResMsg.Code)
		os.Exit(0)
	} else {
		fmt.Println("注册失败...")
		fmt.Println("code:", registerResMsg.Code)
		fmt.Println("err_info:", registerResMsg.ErrorInfo)
		os.Exit(0)
	}

	return
}

func (this *UserProcess)Login(userId int, userPwd string) (err error) {
	fmt.Println("进入登录系统...")
	fmt.Println("用户ID：", userId)
	fmt.Println("用户密码：", userPwd)

	// 1.连接到服务器
	conn , err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		fmt.Println("net.Dial error")
	}
	defer conn.Close()
	fmt.Printf("Login conn addr %p", conn)
	// 2. 通过conn发送消息
	// 创建Message结构体对象
	var msg message.Message
	msg.Type = message.LoginMsgType
	// 创建LoginMsg结构体对象
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	// 把LoginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marsha loginMsg error")
		return
	}
	msg.Data = string(data)

	// 把Message序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marsha Message error")
		return
	}

	// 发送data到服务器
	transfer := &utils.Transfer{
		Conn:conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg error:", err)
		return
	}

	// 3. 获取服务器返回内容
	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) error")
		return
	}

	var loginResMsg message.LoginResMsg

	// 反序列化message
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	if loginResMsg.Code  == 200 {
		currentUser.Conn = conn
		currentUser.UserId = userId
		currentUser.UserStatus = message.UserOnline

		fmt.Printf("**** id为 %d 的用户登录成功 ****\n", userId)
		fmt.Println("code:", loginResMsg.Code)
		fmt.Printf("已在线用户列表id:")
		for _, id := range loginResMsg.UsersId {
			if id == userId {
				continue
			}
			fmt.Printf(" %d", id)
			// 把在线用户保存到map
			user := &message.User{
				UserId:id,
				UserStatus:message.UserOnline,
			}
			onlineUsers[id] = user
		}
		println()
		go serverProcessMes(conn)

		// 进入菜单
		for {
			ShowMenu()
		}

	} else {
		fmt.Println("登录失败...")
		fmt.Println("code:", loginResMsg.Code)
		fmt.Println("err_info:", loginResMsg.ErrorInfo)
	}

	return
}

func (this *UserProcess)Logout() (err error) {
	fmt.Println("进入退出系统...")
	fmt.Println("用户Id：", currentUser.UserId)

	// 创建Message结构体对象
	var msg message.Message
	msg.Type = message.LogoutMsgType
	// 创建LoginMsg结构体对象
	var logoutMsg message.LogoutMsg
	logoutMsg.UserId = currentUser.UserId

	// 把LoginMsg序列化
	data, err := json.Marshal(logoutMsg)
	if err != nil {
		fmt.Println("json.Marsha loginMsg error")
		return
	}
	msg.Data = string(data)

	// 把Message序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marsha Message error")
		return
	}

	// 发送data到服务器
	transfer := &utils.Transfer{
		Conn:currentUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg error:", err)
		return
	}
	return
}