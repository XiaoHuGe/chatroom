package services

import "fmt"

const maxOnlineUser = 100

var (
	userMgr *UserMgr
)

type UserMgr struct {
	users map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		users:make(map[int]*UserProcess, maxOnlineUser),
	}
}

func (this *UserMgr)AddOnlineUser(userProcess *UserProcess) {
	this.users[userProcess.userId] = userProcess
}

func (this *UserMgr)UpdateOnlineUser(userProcess *UserProcess) {
	this.AddOnlineUser(userProcess)
}

func (this *UserMgr)DeleteOnlineUser(userProcess *UserProcess) {
	delete(this.users, userProcess.userId)
}

func (this *UserMgr)QueryOnlineUser(userProcess *UserProcess) (up *UserProcess) {
	up, ok := this.users[userProcess.userId]
	if !ok {
		fmt.Printf("查找的用户-%d-不在线", userProcess.userId)
		return
	}
	return
}

// 返回所有在线用户
func (this *UserMgr)AllOnlineUser() (map[int]*UserProcess) {
	return this.users
}