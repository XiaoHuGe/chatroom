package services

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

func (this *SmsProcess)SendGroupMsg(msg *message.Message) (err error) {

	// 创建smsMsg结构体
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 遍历用户
	for id, up := range userMgr.users{
		if id == smsMsg.UserId {
			continue
		}
		if up.UserStatus == message.UserOffline {
			fmt.Printf("id为【%d的】用户不在线...请留言【功能未实现】\n", id)
			continue
		}
		// 发送data到服务器
		transfer := &utils.Transfer{
			Conn:up.Conn,
		}
		err = transfer.WritePkg(data)
		if err != nil {
			fmt.Println("utils.WritePkg error:", err)
			return
		}
	}

	return
}

func (this *SmsProcess)SendPrivateChatMsg(msg *message.Message) (err error){
	// 创建smsMsg结构体
	var privateChatSmsMsg message.PrivateChatSmsMsg
	err = json.Unmarshal([]byte(msg.Data), &privateChatSmsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 查看是否在线
	up, ok := userMgr.users[privateChatSmsMsg.DestUserId]
	if !ok {
		fmt.Println("用户不存在")
		return
	}
	if up.UserStatus == message.UserOffline {
		fmt.Println("用户不在线...请留言【功能未实现】")
		return
	}

	// 发送data到服务器
	transfer := &utils.Transfer{
		Conn:up.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg error:", err)
		return
	}

	return
}