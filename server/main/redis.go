package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func InitRedis(address string, maxIdle, maxActive int, idleTimeout time.Duration)  {
	pool = &redis.Pool{
		MaxIdle: maxIdle, // 最大空闲链接数
		MaxActive: maxActive, // 数据库最大连接数  0：没有限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
	}
}
