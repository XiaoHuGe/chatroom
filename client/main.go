package main

import "fmt"

var (
	userId int
	userPwd string
)

func main() {
	fmt.Println("--------欢迎登录聊天系统--------")
	fmt.Println("1.登录聊天系统")
	fmt.Println("2.注册用户")
	fmt.Println("3.退出系统")
	fmt.Println("请选择1-3")
	var num int
	fmt.Scanln(&num)
	for {
		switch num {
			case 1:
				fmt.Println("--------登录聊天系统-------")
				fmt.Println("请输入用户ID：")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码：")
				fmt.Scanln(&userPwd)
				err := login(userId, userPwd)
				if err != nil {
					println("登录失败")
				}
				return
			case 2:
				fmt.Println("--------注册聊天系统--------")
				return
			case 3:
				fmt.Println("--------退出聊天系统--------")
				return
			default:
				fmt.Println("输入错误，重新输入1-3")
				fmt.Scanln(&num)
				break

		}
	}
}