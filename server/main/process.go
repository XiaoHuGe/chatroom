package main

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"chatroom/server/services"
	"fmt"
	"io"
	"net"
)

type Process struct {
	Conn net.Conn
}

func (this *Process)ServerProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginMsgType:
			 userPro := &services.UserProcess{
				Conn:this.Conn,
			 }
			 userPro.ServerProcessLogin(msg)
		case message.RegisterMsgType:
			userPro := &services.UserProcess{
				Conn:this.Conn,
			}
			userPro.ServerProcessRegister(msg)
		case message.LogoutMsgType:
			userPro := &services.UserProcess{
				Conn:this.Conn,
			}
			userPro.ServerProcessLogout(msg)
		default:
			println("消息类型不存在")
	}
	return
}


func (this *Process)process() (err error) {

	for {
		// 读取客户端消息
		transfer := &utils.Transfer{
			Conn:this.Conn,
		}
		msg, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出....")
				return err
			} else {
				fmt.Println("readPkg error: ", err)
				return err
			}
		}
		fmt.Println("msg :", msg)

		// 处理客户端消息
		this.ServerProcessMsg(&msg)
	}
	return
}
