package model
import(
	"net"
	"go_code/hello_project/chapter16chatroom/common/message"

)




type CurUser struct{
	Conn net.Conn
	message.User
}