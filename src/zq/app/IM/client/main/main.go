package main

import (
	"fmt"
	"zq/app/IM/client/process"
)

var key int
var userId int
var userPwd string

func main() {
	for true {
		fmt.Println("1-用户登录")
		fmt.Scan(&key)

		switch key {
		case 1:

			//登录
			fmt.Println("用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("密码")
			fmt.Scanf("%s\n", &userPwd)

			// 测试数据 userId: 100  userPwd: pwd
			// 测试数据 userId: 101  userPwd: pwd

			userProcess := process.UserProcess{}
			// 用户登录
			userProcess.Login(userId, userPwd)

		}
	}

	// 注册
	//userProcee.Register(101, "zongqi_101", "pwd")

	//userProcee.Login(101, "pwd")
}
