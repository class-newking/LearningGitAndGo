package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

var ch = make(chan int, 2)

func CreateConn(num int) {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("绑定端口错误：%s\n", err)
		return
	}
	defer conn.Close()
	str := "hello, i am" + strconv.Itoa(num)
	for {
		time.Sleep(time.Second*5)
		_,err := conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("写入错误：%s\n", err.Error())
		} else {
			fmt.Printf("发送：%s\n", str)
		}

		buff := make([]byte, 128)
		datalen, err := conn.Read(buff)
		if err == io.EOF {
			fmt.Println("服务器关闭")
			ch <- 1
			return
		}
		if err == nil && datalen > 0 {
			fmt.Printf("接收：%s\n", string(buff[:datalen]))
		}
	}
}

func main() {
	for i := 1; i <= 2; i++ {
		go CreateConn(i)
	}
	<- ch
	<- ch
}