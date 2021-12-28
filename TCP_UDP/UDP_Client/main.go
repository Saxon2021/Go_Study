package main

import (
	"fmt"
	"net"
)

// UDP 客户端
func main() {

	// 连接到UDP服务器
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})

	// 连接失败
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}

	// 关闭Socket
	defer socket.Close()

	// 要发送的数据
	sendData := []byte("Hello server")

	// 向服务端发送数据
	_, err = socket.Write(sendData) // 发送数据

	// 发送数据失败
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}

	// 创建一个接收数据的切片
	data := make([]byte, 4096)

	// 读取服务端发送的数据
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据

	// 读取服务端的数据失败
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}

	// 打印来自服务端的数据
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
