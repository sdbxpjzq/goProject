package process

import (
	"encoding/json"
	"fmt"
	"net"
	"zq/app/IM/client/utils"
	"zq/app/IM/common/message"
)

func ServerProcessMes(conn net.Conn) {

	transfer := utils.Transfer{
		Conn: conn,
	}
	// 客户端不停的读取 服务端信息
	for {
		fmt.Println("客户端正在等待读取服务端发送的消息")
		mes, err := transfer.ReadPkg()
		if err != nil {
			println("serverProcessMes", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStateMes *message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStateMes)
			updateUserStatus(notifyUserStateMes)
		}

	}
}
