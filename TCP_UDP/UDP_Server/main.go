package main

import (
	"fmt"
	"net"
)

// UDP/server/main.go

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	// 注册关闭监听事件
	defer listen.Close()

	for {

		// 创建一个切片
		data := make([]byte, 1024)

		// 接收数据
		n, addr, err := listen.ReadFromUDP(data)

		// 接收数据失败
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}

		// 打印接收到的数据
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		// 向客户端发送数据
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据

		// 发送数据失败
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
