package main

import (
	"flag"
	"fmt"
	"net"
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

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			fmt.Println("公聊模式")
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("更新用户名")
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
	fmt.Println("connect success")

	client.Run()
}
