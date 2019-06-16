package process

import (
	"fmt"
	"ysqi/app/IM/common/message"
)

var onLineUsers = make(map[int]*message.User, 100)

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onLineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onLineUsers[notifyUserStatusMes.UserId] = user
	outPutOnlineUser()
}

func outPutOnlineUser() {
	for key, value := range onLineUsers {
		fmt.Println(key, value)
	}
}
