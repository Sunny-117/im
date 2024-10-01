package process
import(
	"fmt"
	"os"
	"go_code/hello_project/chapter16chatroom/client/utils"
	"net"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/json"

)

//显示登陆成功后的界面...
func ShowMenu(){
	fmt.Println("-----------登陆成功-------------------")
	fmt.Println("-----------1. 显示在线用户列表---------")
	fmt.Println("-----------2. 发送消息----------------")
	fmt.Println("-----------3. 信息列表----------------")
	fmt.Println("-----------4. 退出系统----------------")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1 : 
			// fmt.Println("显示在线用户列表~")
			outputOnlineUsers()
		case 2 : 
			//fmt.Println("发送消息~")
			fmt.Println("请输入你要发送的消息：")
			fmt.Scanf("%s\n",&content)
			//向服务器发送消息
			smsProcess.SendGroupMes(content)
		case 3 : 
			fmt.Println("显示在线用户列表~")
		case 4 : 
			fmt.Println("选择退出系统~")
			os.Exit(0)
		default :
			fmt.Println("你输入的选项不正确...")
	}
}


//和服务器保持通讯的协程
func serverProcessMes(conn net.Conn){
	tf := &utils.Transfer{
		Conn : conn,
	}

	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes , err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg error = ", err)
			return
		}
		//根据消息类型去处理
		switch mes.Type{
			case message.NotifyUserStatusMesType : //有人上线了
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				//更新map
				UpdateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType : //有人群发消息
				outputGroupMes(&mes)
			default :
				fmt.Println("服务器返回未知消息类型")
		}
	}
}