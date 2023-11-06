package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	// create client
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	// connect server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.conn = conn

	return client
}

// handle message from server
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn) // 永久阻塞监听，一旦client.conn有数据，就直接拷贝到Stdout标准输出上
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true

	} else {
		fmt.Println("please input valid number")
		return false
	}

}

func (client *Client) PublicChat() {
	var chatMsg string
	fmt.Println("please input your message, exit.")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		// message is not empty
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error: ", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("please input your message, exit.")
		fmt.Scanln(&chatMsg)
	}

}

func (client *Client) UpdateName() bool {
	fmt.Println("please input your name")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error: ", err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "server ip address(default is 127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "server port(default is 8888)")
}
func main() {
	// command line resolve
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("fail to connect client")
		return
	}

	// 单独开启一个goroutine，处理server的消息
	go client.DealResponse()

	fmt.Println("connect success")

	client.Run()
}
