package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工程模式返回UserDao
func NewUserDao(pool *redis.Pool) (userDao *UserDao)  {
	userDao = &UserDao{
		pool:pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 从redis查询用户id是否存在
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误!
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应id
			err = ERROR_UESR_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err :", err)
		return
	}

	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 从连接池获取一个连接
	conn := this.pool.Get()
	defer  conn.Close()
	// 查询用户id
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 判断密码
	if user.UserPwd != userPwd {
		err = ERROR_UESR_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	// 从连接池获取一个连接
	conn := this.pool.Get()
	defer  conn.Close()
	// 查询用户id
	_, err = this.getUserById(conn, user.UserId)
	fmt.Println("注册 UserId: ", user.UserId)
	// 没有错误代表之前注册过
	if err == nil {
		err = ERROR_UESR_EXISTS
		return
	}

	data, err :=  json.Marshal(user) // 序列化user

	// 保存用户信息到redis
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户失败", err)
		return
	}
	return
}
