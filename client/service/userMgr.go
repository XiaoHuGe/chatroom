package service

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

const maxOnlineUser = 100

var users map[int]*message.User = make(map[int]*message.User, maxOnlineUser)
var currentUser model.CurrentUser
func GetAllUser() {
	fmt.Printf("当前在线用户列表id:")
	for id, user := range users{
		if user.UserStatus == message.UserOnline {
			fmt.Printf(" %d", id)
		}
	}
	fmt.Println()
}

// 更新map
func UpdateUserStatus(notifyUserStatus *message.NotifyUserStatusMsg) {
	user, ok := users[notifyUserStatus.UserId]
	if !ok {
		// 把上线用户添加到map
		user = &message.User{
			UserId:notifyUserStatus.UserId,
			//UserStatus:notifyOnlineMsg.UserStatus,
		}
	}
	user.UserStatus = notifyUserStatus.UserStatus
	users[notifyUserStatus.UserId] = user
	GetAllUser()
}