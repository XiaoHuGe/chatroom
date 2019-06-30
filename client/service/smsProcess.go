package service

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

func (this *SmsProcess)SendGroupMsg(content string) (err error) {
	// 创建Message结构体对象
	var msg message.Message
	msg.Type = message.SmsMsgType
	// 创建LoginMsg结构体对象
	var smsMsg message.SmsMsg
	smsMsg.UserId = currentUser.UserId
	smsMsg.Content = content
	// 把LoginMsg序列化
	data, err := json.Marshal(smsMsg)
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
	//err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("utils.WritePkg error:", err)
		return
	}
	return
}