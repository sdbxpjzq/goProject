package lib

import (
	"github.com/gomodule/redigo/redis"
)

func InitPool() (redisPool *redis.Pool) {
	// redis 的时候用
	redisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) { // 初始化链接
			return redis.Dial("tcp", "47.244.160.71:6379")
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         10, //最大空闲链接数
		MaxActive:       0,  // 和数据库的最大链接数, 0表示没限制
		IdleTimeout:     0,  // 最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}
	return
}
