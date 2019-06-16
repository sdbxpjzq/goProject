package main

import (
	"net"
	"ysqi/app/IM/common/message"
	process2 "ysqi/app/IM/server/process"
	utils2 "ysqi/app/IM/server/utils"
)

type Processor struct {
	Conn net.Conn
}

/**
根据消息的种类, 调用不同函数处理不同的消息
*/
func (p *Processor) serverProcessMes(mes *message.Message) {
	switch mes.Type {
	case message.LoginMesType:
		// 登录
		userProcess := process2.UserProcess{
			Conn: p.Conn,
		}
		userProcess.ServerProcessLoginMes(mes)

	case message.RegisterMesType:
		// 注册
		userProcess := process2.UserProcess{
			Conn: p.Conn,
		}
		userProcess.ServerProcessRegisterMes(mes)

	}

}

/**
入口层
*/
func (p *Processor) ProcessEnter() (err error) {
	// 循环读取客户端发送的消息
	for {
		// 读取数据
		utils := utils2.Transfer{
			Conn: p.Conn,
		}
		mes, err := utils.ReadPkg()
		if err != nil {
			return err
		}
		p.serverProcessMes(&mes)
	}
}
