package process

import (
	"encoding/json"
	"fmt"
	"net"
	"ysqi/app/IM/client/utils"
	"ysqi/app/IM/common/message"
)

type UserProcess struct {
	Conn net.Conn
}

func connectRedis() (conn net.Conn, err error) {
	conn, error := net.Dial("tcp", "127.0.0.1:8889")
	if error != nil {
		fmt.Println("连接错误", error)
		return
	}
	return
}

func (u *UserProcess) Register(userId int, userName string, userPwd string) {
	conn, err := connectRedis()
	if err != nil {
		fmt.Println("连接错误", err)
		return
	}
	defer conn.Close()
	registerMes := message.RegisterMes{
		UserId:   userId,
		UserName: userName,
		UserPwd:  userPwd,
	}
	registerBytes, error := json.Marshal(registerMes)
	if error != nil {
		fmt.Println("registerMes json错误", error)
		return
	}
	mes := message.Message{
		Type: message.RegisterMesType,
		Data: string(registerBytes),
	}
	// 序列化mes
	bytes, error := json.Marshal(mes)
	if error != nil {
		fmt.Println("mes json错误", error)
		return
	}

	// 接收服务端返回的消息
	transfer := utils.Transfer{
		Conn: conn,
	}

	// 写bytes
	error = transfer.WritePkg(bytes)
	if error != nil {
		fmt.Println("WritePkg err", error)
		return
	}

	mes, error = transfer.ReadPkg()

	var registerResMes message.RegisterResMes
	json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return

}

func (u *UserProcess) Login(userId int, userPwd string) {

	conn, error := net.Dial("tcp", "127.0.0.1:8889")
	if error != nil {
		fmt.Println("连接错误", error)
		return
	}
	defer conn.Close()

	// 创建loginMes
	loginMes := message.LoginMes{
		UserId:   userId,
		UserName: "zongqi",
		UserPwd:  userPwd,
	}

	// 序列化loginMes
	LoginBytes, error := json.Marshal(loginMes)

	if error != nil {
		fmt.Println("loginMes json错误", error)
		return
	}

	mes := message.Message{
		Type: message.LoginMesType,
		Data: string(LoginBytes),
	}
	// 序列化mes
	bytes, error := json.Marshal(mes)
	if error != nil {
		fmt.Println("mes json错误", error)
		return
	}

	// 接收服务端返回的消息
	transfer := utils.Transfer{
		Conn: conn,
	}

	// 写bytes
	error = transfer.WritePkg(bytes)
	if error != nil {
		fmt.Println("WritePkg err", error)
		return
	}

	mes, error = transfer.ReadPkg()
	if error != nil {
		return
	}

	var loginResMes message.LoginResMes
	json.Unmarshal([]byte(mes.Data), &loginResMes)
	fmt.Println(loginResMes)
	if loginResMes.Code == 200 {
		// 登录成功
		fmt.Println("登录成功")
		for _, v := range loginResMes.UserIds {
			//if v == userId {
			//	continue
			//}
			user := message.User{
				UserId:     userId,
				UserStatus: message.UserOnline,
			}
			onLineUsers[v] = &user
		}

		go ServerProcessMes(conn)
	} else {
		fmt.Println(loginResMes.Error)
	}
}
