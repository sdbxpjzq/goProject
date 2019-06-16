package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"ysqi/app/IM/common/message"
)

type Transfer struct {
	// 属性
	Conn net.Conn
	Buf  [8096]byte
}

/**
读取message
*/
func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := t.Buf
	_, err = t.Conn.Read(buf[:4])
	if err != nil {
		fmt.Println("数据读取失败3", err)
		return
	}

	// buf[:4] 转换uint32的转换
	pkgLen := binary.BigEndian.Uint32(buf[:4])
	// 根据pkgLen读取消息内容
	n, err := t.Conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("数据读取失败4", err)
		return
	}
	// 反序列化, & 结构体是值传递
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("数据反序列失败", err)
		return
	}
	return
}

/**
写message
*/
func (t *Transfer) WritePkg(data []byte) error {
	// 7 bytes 就是要发送的数据
	// 7.1 先把 bytes 的长度发送给服务器-- 现获取 bytes 的长度, --> 将长度转换成 切片, 使用 binary包

	// 发送data长度
	var bytesLen [4]byte
	binary.BigEndian.PutUint32(bytesLen[0:4], uint32(len(data)))
	// 发送长度
	_, error := t.Conn.Write(bytesLen[0:4])
	if error != nil {
		fmt.Println("data长度发送失败")
		return error
	}

	// 发送data消息
	_, error = t.Conn.Write(data)
	if error != nil {
		fmt.Println("data发送失败")
		return error
	}
	return nil

}

//
//var bytesLen [4]byte
//binary.BigEndian.PutUint32(bytesLen[0:4], uint32(len(bytes)))
//// 发送长度
//_, error = conn.Write(bytesLen[0:4])
//if error != nil {
//fmt.Println("长度发送失败", error)
//return
//}
//
//// 发送消息数据
//_, error = conn.Write(bytes)
//if error != nil {
//fmt.Println("数据体发送失败", error)
//return
//}
