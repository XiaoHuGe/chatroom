package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//这里调用总控, 创建一个
	pro := &Process{
		Conn : conn,
	}
	err := pro.process()
	if err != nil {
		fmt.Println("pro.process err:", err)
		return
	}
}

func main() {

	fmt.Println("服务器在8899端口监听...")
	listener, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		fmt.Println("net.Listen error: ", err)
	}
	defer listener.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error: ", err)
		}
		// 连接成功，开启一个协程保持与客户端的连接
		go process(conn)
	}
}