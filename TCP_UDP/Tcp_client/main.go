package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// 与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20001")

	// 连接服务端失败
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	// 从标准输入中获取内容
	inputReader := bufio.NewReader(os.Stdin)

	for {

		// 读取用户输入的内容
		input, _ := inputReader.ReadString('\n')

		// 去除用户输入的结尾
		inputInfo := strings.Trim(input, "\r\n")

		// 用户输入为Q 则断开与服务端的连接
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 向服务端发送数据
		_, err = conn.Write([]byte(inputInfo))

		// 发送数据失败
		if err != nil {
			return
		}

		buf := make([]byte, 512)

		// 读取服务端返回的数据
		n, err := conn.Read(buf)

		// 读取失败
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
