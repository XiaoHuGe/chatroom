package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func ReadPkg(conn net.Conn) (msg message.Message, err error){

	buff := make([]byte, 8096)

	// 获取客户端发送的长度
	_, err = conn.Read(buff[:4])
	if err != nil {
		return
	}

	// []byte转Uint32
	var pkgLen = binary.BigEndian.Uint32(buff[:4])
	// 获取数据
	n, err := conn.Read(buff[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	// 反序列化数据
	err = json.Unmarshal(buff[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送data长度.把长度转化成[]byte
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n ,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		//println("conn.Write(buf[:4]) error")
		return
	}

	fmt.Printf("data len:%d, data:%s", len(data), string(data))
	fmt.Println()

	// 发送data
	_, err = conn.Write(data)
	if err != nil {
		//println("conn.Write(data) error")
		return
	}

	return
}