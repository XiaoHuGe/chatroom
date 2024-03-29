package service

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu()  {
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送群消息---------")
	fmt.Println("-------3. 发送点对点消息---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")

	var key int
	var destUserId int
	var content string
	sp := SmsProcess{}
	up := &UserProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
		case 1:
			GetAllUser()
		case 2:
			fmt.Println("输入要群发的消息：")
			fmt.Scanf("%s\n", &content)
			sp.SendGroupMsg(content)
		case 3:
			fmt.Println("输入好友id：")
			fmt.Scanf("%d\n", &destUserId)
			fmt.Println("输入要发送的消息：")
			fmt.Scanf("%s\n", &content)
			sp.SendProriveChatMsg(destUserId, content)
		case 4:
			fmt.Println("退出系统")
			up.Logout()
			//os.Exit(0)
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
				switch notifyUserStatusMsg.UserStatus {
					case message.UserOnline:
						fmt.Printf("用户id为 %d 的用户上线了...\n", notifyUserStatusMsg.UserId)
					case message.UserOffline:
						fmt.Printf("用户id为 %d 的用户退出了...\n", notifyUserStatusMsg.UserId)
				}
				// 添加上线用户到map
				UpdateUserStatus(&notifyUserStatusMsg)
			case message.LogoutResMsgType:
				var logoutResMsg message.LogoutResMsg
				// 反序列化message
				err = json.Unmarshal([]byte(msg.Data), &logoutResMsg)
				if err != nil {
					fmt.Println("json.Unmarshal error:", err)
					return
				}
				fmt.Println("logoutResMsg Code:", logoutResMsg.Code)
				if logoutResMsg.Code == 200 {
					fmt.Printf("退出成功...\n")
					fmt.Println("code:", logoutResMsg.Code)
					os.Exit(0)
				} else {
					fmt.Println("退出失败...检查网络连接")
					fmt.Println("code:", logoutResMsg.Code)
					fmt.Println("err_info:", logoutResMsg.ErrorInfo)
					return
				}
			case message.SmsMsgType:
				var smsMsg message.SmsMsg
				// 反序列化message
				err = json.Unmarshal([]byte(msg.Data), &smsMsg)
				if err != nil {
					fmt.Println("json.Unmarshal error:", err)
					return
				}
				info := fmt.Sprintf("接收到用户id为【%d】 的群消息，内容为【 %s】", smsMsg.UserId, smsMsg.Content)
				fmt.Println(info)
			case message.PrivateChatSmsMsgType:
				var privateChatSmsMsg message.PrivateChatSmsMsg
				err = json.Unmarshal([]byte(msg.Data), &privateChatSmsMsg)
				if err != nil {
					fmt.Println("json.Unmarshal error:", err)
					return
				}
				info := fmt.Sprintf("接收到用户id为【%d】 的消息，内容为【 %s】", privateChatSmsMsg.UserId, privateChatSmsMsg.Content)
				fmt.Println(info)
		}
	}
}