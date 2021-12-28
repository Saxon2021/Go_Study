package main

import (
	"bufio"
	"fmt"
	"net"
)

// 定义处理程序
func process(conn net.Conn) {
	// 关闭远程连接
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)

		// var buf [128]byte

		buf := make([]byte, 128)

		// 读取数据
		n, err := reader.Read(buf)

		// 读取客户端
		if err != nil {
			fmt.Println("read from clien failed, err:", err)
			break
		}
		recvInfo := string(buf[:n])

		fmt.Println("Client Data: ", recvInfo)

		conn.Write([]byte("收到"))
	}
}
func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:20001")

	if err != nil {
		fmt.Println("Listen failed, err:", err)
		return
	}
	for {
		// 与客户端建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}
}
