package service

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {

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
	transfer := utils.Transfer{
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

	var loginResMsg message.LoginResMsg

	// 反序列化message
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	if loginResMsg.Code  == 200 {
		fmt.Println("登录成功...")
		fmt.Println("code:", loginResMsg.Code)
	} else {
		fmt.Println("登录失败...")
		fmt.Println("code:", loginResMsg.Code)
		fmt.Println("err_info:", loginResMsg.ErrorInfo)
	}

	return
}