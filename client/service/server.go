package service

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu(userId int)  {
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
		case 1:
			GetAllUser()
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
	}
}

// 监听服务器发回的消息
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg:", err)
			return
		}
		switch msg.Type {
			case message.NotifyUserStatusMsgType:
				// 创建NotifyOnlineMsg结构体
				var notifyUserStatusMsg message.NotifyUserStatusMsg
				err = json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
				if err != nil {
					fmt.Println("json.Unmarshal err: ", err)
					return
				}
				fmt.Printf("用户id为 %d 的用户上线了...\n", notifyUserStatusMsg.UserId)
				// 添加上线用户到map
				UpdateUserStatus(&notifyUserStatusMsg)
		}
	}
}