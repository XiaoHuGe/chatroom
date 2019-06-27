package model

import (
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