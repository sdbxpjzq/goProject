package process

import (
	"errors"
)

type UserMgr struct {
	onlineUserIds map[int]*UserProcess
}

var userMgr *UserMgr

// 初始化
func init() {
	userMgr = &UserMgr{
		onlineUserIds: make(map[int]*UserProcess, 1024),
	}
}

// 增加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUserIds[up.UserId] = up
}

// 删除
func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUserIds, userId)
}

// 查询
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUserIds
}

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUserIds[userId]
	if !ok {
		err = errors.New("词用户ID不在线")
		return
	}
	return
}
