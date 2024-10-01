package process
import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"go_code/hello_project/chapter16chatroom/client/model"

)

//客户端维护的在线用户map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10) 
var curUser model.CurUser //负责聊天的变量，在用户登录成功后，完成对它的初始化

//更新接收到的通知信息
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	
	user ,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //新用户
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
		onlineUsers[notifyUserStatusMes.UserId] = user
	}
	user.UserStatus = notifyUserStatusMes.Status
	outputOnlineUsers()
	
}

//显示当前在线用户
func outputOnlineUsers(){
	fmt.Println("当前其它在线用户：")
	fmt.Println("=========================================================")
	for id,user := range onlineUsers{
		fmt.Println("用户id：\t",id," , 用户名:\t",user.UserName)
	}
	fmt.Println("=========================================================")
}