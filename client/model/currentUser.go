package model

import (
	"chatroom/common/message"
	"net"
)

type CurrentUser struct {
	Conn net.Conn
	message.User
}
