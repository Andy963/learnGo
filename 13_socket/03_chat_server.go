package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func writeMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func makeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": online"
	return buf
}

// 处理用户连接
func handleCon(conn net.con) {
	defer conn.Close()

	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	// 创建一个结构体，默认，用户名为网络地址
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli

	// 新开一个协程，专门给当前客户端发送信息
	go writeMsgToClient(cli, conn)

	// 广播在线用户
	//message <- "[" + cli.Addr + "]" + cli.Name + ": online"
	message <= makeMsg(cli, 'online')

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("conn.Read err:", err)
				return
			}
			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				// 遍历map, 给当前用户 发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else {
				message <- makeMsg(cli, msg)
			}
		}
	}

	for {

	}

}

func manager() {
	// 给map 分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message // 没有消息前，阻塞
		// 遍历，给map每个成员都 发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net listen  err", err)
		return
	}

	// 新开一个协程，只要有消息来了，给每一个成员发送消息

	defer listener.Close()
	// 循环阻塞，等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err", err)
			continue
		}

		go handleCon(conn)
	}
}
