package main

import (
	"fmt"
	"net"
	"ysqi/app/IM/server/model"
	"ysqi/lib"
)

func process(conn net.Conn) {
	defer conn.Close()

	// 调用控制入口
	enter := Processor{
		Conn: conn,
	}
	err := enter.ProcessEnter()
	if err != nil {
		fmt.Println("utils.ReadPkg失败", err)
		return
	}

}

func init() {
	// 初始化redis连接池
	redisPool := lib.InitPool()
	model.MyUserDao = model.NewUserDao(redisPool)
}

func main() {
	// 监听端口
	listener, error := net.Listen("tcp", "0.0.0.0:8889")
	defer listener.Close()

	if error != nil {
		fmt.Println("监听错误", error)
		return
	}

	// 监听成功
	for {
		conn, error := listener.Accept()
		if error != nil {
			fmt.Println("连接失败")
			// 这里不需要return, 可能只有一个链接失败
		}

		// 连接成功, 启动协程, 和客户端保持通讯

		go process(conn)

	}

}
