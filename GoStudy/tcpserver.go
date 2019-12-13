package main

import (
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	ip := conn.RemoteAddr().String()
	fmt.Printf("客户端：%s\n", ip)
	for {
		buff := make([]byte, 128)
		bufflen, err := conn.Read(buff)
		if err == io.EOF {
			fmt.Printf("%s 关闭\n", ip)
			return
		}
		if err != nil {
			fmt.Printf("%s 读取错误：%s\n", ip, err.Error())
			continue
		}
		if bufflen > 0 {
			fmt.Printf("%s 接收到的信息：%s\n", ip, string(buff[0:bufflen]))
			conn.Write(buff[0:bufflen])
		}
	}
}

func main () {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("监听失败，错误：", err)
		return
	}
	defer listen.Close()
	fmt.Println("tcp 监听 127.0.0.1:9999")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败，错误：", err)
			continue
		}
		go handle(conn)
	}
}
