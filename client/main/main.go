package main

import (
	"chatroom/client/service"
	"fmt"
	"os"
)

var (
	userId int
	userPwd string
	userName string
)

func main() {
	for {
		fmt.Println("--------欢迎登录聊天系统--------")
		fmt.Println("1.登录聊天系统")
		fmt.Println("2.注册用户")
		fmt.Println("3.退出系统")
		fmt.Println("请选择1-3")

		var num int
		fmt.Scanln(&num)
		switch num {
			case 1:
				fmt.Println("--------登录聊天系统-------")
				fmt.Println("请输入用户Id：")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码：")
				fmt.Scanln(&userPwd)
				up := &service.UserProcess{}
				err := up.Login(userId, userPwd)
				if err != nil {
					println("登录失败")
				}
				//return
			case 2:
				fmt.Println("--------注册聊天系统--------")
				fmt.Println("请输入用户Id：")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码：")
				fmt.Scanln(&userPwd)
				fmt.Println("请输入用户昵称：")
				fmt.Scanln(&userName)
				up := &service.UserProcess{}
				err := up.Register(userId, userPwd, userName)
				if err != nil {
					println("登录失败")
				}
			case 3:
				fmt.Println("--------退出聊天系统--------")
				os.Exit(0)
			default:
				fmt.Println("输入错误，重新输入1-3")

		}
	}
}