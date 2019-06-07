package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"ysqi/app/IM/common/message"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式, 创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (u *UserDao) getUserById(conn redis.Conn, id int) (user message.User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			// 没有找到
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	json.Unmarshal([]byte(res), &user)
	return
}

/**
登录
*/

func (u *UserDao) Login(userId int, userPwd string) (user message.User, err error) {
	// 从连接池取链接
	conn := u.pool.Get()
	defer conn.Close()

	user, err = u.getUserById(conn, userId)

	if err != nil {
		// 用户不存在
		return
	}
	if user.UserPwd != userPwd {
		// 密码错误
		err = ERROR_USER_PWD
		return
	}
	return

}

func (u *UserDao) Register(user *message.RegisterMes) (err error) {
	// 从连接池取链接
	conn := u.pool.Get()
	defer conn.Close()
	bytes, _ := json.Marshal(user)
	_, err = conn.Do("hset", "users", user.UserId, string(bytes))
	if err != nil {
		fmt.Println("保存注册用户出错")
		return
	}
	return
}
