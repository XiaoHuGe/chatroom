package service

import (
	"chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
func GetAllUser() {
	fmt.Printf("当前在线用户列表id:")
	for id, _ := range onlineUsers{
		fmt.Printf(" %d", id)
	}
	fmt.Println()
}

func UpdateUserStatus(notifyUserStatus *message.NotifyUserStatusMsg) {
	// 把上线用户添加到map
	user, ok := onlineUsers[notifyUserStatus.UserId]
	if !ok {
		user = &message.User{
			UserId:notifyUserStatus.UserId,
			//UserStatus:notifyOnlineMsg.UserStatus,
		}
	}
	user.UserStatus = notifyUserStatus.UserStatus
	onlineUsers[notifyUserStatus.UserId] = user
}