package message

const LoginMesType = "LoginMes"
const LoginResMesType = "LoginResMes"

const RegisterMesType = "RegisterMes"
const RegisterResMesType = "RegisterResMes"
const NotifyUserStatusMesType = "NotifyUserStatusMes"

const (
	UserOnline = iota
	UserOffline
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 登录消息
type LoginMes struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code    int    `json:"code"`     // 500 用户没有注册, 200 登录成功
	Error   string `json:"error"`    // 错误信息
	UserIds []int  `json:"user_ids"` // 保存用户id切片
}

// 注册消息
type RegisterMes struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 500 注册失败, 200 注册成功
	Error string `json:"error"` // 错误信息
}

// 用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}
