package services

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

func (this *SmsProcess)SendGroupMsg(msg *message.Message) (err error){

	// 创建LoginMsg结构体
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
	for id, up := range userMgr.onlineUser{
		if id == smsMsg.UserId {
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