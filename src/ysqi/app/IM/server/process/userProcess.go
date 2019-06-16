package process

import (
	"encoding/json"
	"fmt"
	"net"
	"ysqi/app/IM/common/message"
	"ysqi/app/IM/server/model"
	"ysqi/app/IM/server/utils"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// userId 要通知其他的在线用户	, 我上线
func (u *UserProcess) NotifyOtherOnlineUser(userId int) {
	for id, _ := range userMgr.onlineUserIds {
		if id == userId {
			// 自己登录不需要通知
			continue
		}
		NotifyUserStatusMes := message.NotifyUserStatusMes{
			UserId: userId,
			Status: message.UserOnline,
		}
		NotifyUserStatusMesBytes, err := json.Marshal(NotifyUserStatusMes)
		if err != nil {
			fmt.Println("NotifyOtherOnlineUser 序列化出错", err)
			return
		}

		mes := message.Message{
			Type: message.NotifyUserStatusMesType,
			Data: string(NotifyUserStatusMesBytes),
		}

		bytes, err := json.Marshal(mes)
		//发送
		// 发送data
		transfer := &utils.Transfer{
			Conn: u.Conn,
		}
		err = transfer.WritePkg(bytes)
		return
	}

}

// 处理注册
func (u *UserProcess) ServerProcessRegisterMes(mes *message.Message) {
	var RegisterMes message.RegisterMes
	err := json.Unmarshal([]byte(mes.Data), &RegisterMes)
	if err != nil {
		fmt.Println("loginMes 反序列化失败")
		return
	}
	// 注册入库
	err = model.MyUserDao.Register(&RegisterMes)
	// 返回消息
	var registerResMes message.RegisterResMes
	if err != nil {
		registerResMes = message.RegisterResMes{
			Code:  500,
			Error: err.Error(),
		}
	} else {
		registerResMes = message.RegisterResMes{
			Code:  200,
			Error: "注册成功",
		}
	}
	registerResMesBytes, err := json.Marshal(registerResMes)
	// resMes
	var resMes = message.Message{
		Type: message.RegisterResMesType,
		Data: string(registerResMesBytes),
	}

	resMesBytes, err := json.Marshal(resMes)
	// 发送data
	transfer := utils.Transfer{
		Conn: u.Conn,
	}
	err = transfer.WritePkg(resMesBytes)
	return

}

/**
处理登录
*/
func (u *UserProcess) ServerProcessLoginMes(mes *message.Message) {
	// 从mes中获取Data, --> LoginMes
	var loginMes message.LoginMes
	err := json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("loginMes 反序列化失败")
		return
	}

	var loginResMes message.LoginResMes
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		// 登录失败
		loginResMes = message.LoginResMes{
			Code:  500,
			Error: err.Error(),
		}
	} else {
		// 登录成功
		loginResMes = message.LoginResMes{
			Code:  200,
			Error: user.UserName + "登录成功",
		}
		// 添加在线用户
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)
		for id := range userMgr.onlineUserIds {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		// 通知其他在线用户 我上线了
		u.NotifyOtherOnlineUser(loginMes.UserId)
	}
	fmt.Println(loginResMes)

	loginResMesBytes, err := json.Marshal(loginResMes)
	// resMes
	var resMes = message.Message{
		Type: message.LoginResMesType,
		Data: string(loginResMesBytes),
	}

	resMesBytes, err := json.Marshal(resMes)
	// 发送data
	transfer := utils.Transfer{
		Conn: u.Conn,
	}
	err = transfer.WritePkg(resMesBytes)
	return

}
