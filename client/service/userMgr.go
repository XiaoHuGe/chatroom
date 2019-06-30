package service

import (
	"chatroom/common/message"
	"fmt"
)

const maxOnlineUser = 100

var onlineUsers map[int]*message.User = make(map[int]*message.User, maxOnlineUser)
func GetAllUser() {
	fmt.Printf("当前在线用户列表id:")
	for id, user := range onlineUsers{
		if user.UserStatus == message.UserOnline {
			fmt.Printf(" %d", id)
		}
	}
	fmt.Println()
}

// 更新map
func UpdateUserStatus(notifyUserStatus *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[notifyUserStatus.UserId]
	if !ok {
		// 把上线用户添加到map
		user = &message.User{
			UserId:notifyUserStatus.UserId,
			//UserStatus:notifyOnlineMsg.UserStatus,
		}
	}
	user.UserStatus = notifyUserStatus.UserStatus
	onlineUsers[notifyUserStatus.UserId] = user
	GetAllUser()
}