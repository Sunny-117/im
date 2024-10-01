package process
import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/json"
	_ "encoding/binary"
	"net"
	_ "time"
	"go_code/hello_project/chapter16chatroom/client/utils"
	"os"
)

type UserProcessor struct{}

func (this *UserProcessor) Register(userId int , 
	userPwd string , userName string) (err error){

		//connect server
		conn , err := net.Dial("tcp","localhost:8889")
		if err != nil {
			fmt.Println("net.Dial error = ", err)
			return
		}
		defer conn.Close()
		
		//定消息类型
		var mes message.Message
		mes.Type = message.RegisterMesType
		//创建结构体
		var registerMes message.RegisterMes 
		registerMes.User.UserId = userId
		registerMes.User.UserPwd = userPwd
		registerMes.User.UserName = userName

		//4序列化
		data , err := json.Marshal(registerMes)
		if err != nil{
			fmt.Println("json.Marshal error = ",err)
			return
		}
		mes.Data = string(data)
		
		data , err = json.Marshal(mes)
		if err != nil {
			fmt.Println("json.Marshal error = ",err)
			return
		}
		tf := &utils.Transfer{
			Conn : conn,
		}
		//发送data
		err = tf.WritePkg(data)
		if err != nil {
			fmt.Println("注册发送信息错误，error = ",err)
			return
		}

		mes , err = tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg errror = ",err)
			return
		}

		//将mes的data部分反序列化成
		var registerResMes message.RegisterResMes
		err = json.Unmarshal([]byte(mes.Data),&registerResMes)

		//注册成功与否，都推出进程
		if registerResMes.Code == 200 {
			fmt.Println("注册成功，请重新登录")
			os.Exit(0)
		}else{
			fmt.Println(registerResMes.Error)
			os.Exit(0)
		}
		
		return

}



func (this *UserProcessor) Login(userId int , userPwd string) (err error){

	//connect server
	conn , err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error = ", err)
		return
	}
	defer conn.Close()
	
	//定义消息
	var mes message.Message
	mes.Type = message.LoginMesType
	//data
	var loginMes message.LoginMes 
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//序列化，data为切片
	data , err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal error = ",err)
		return
	}

	//这里的mes其实设计得不够好，应把mes理解成res
	mes.Data = string(data)

	data , err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error = ",err)
		return
	}

	//创建实例
	tf := &utils.Transfer{
		Conn : conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg error = ", err)
		return
	}
	fmt.Println("客户端发送报文为：",string(data))
	
	//接收登陆结果
	mes , err = tf.ReadPkg()
	if err != nil{
		fmt.Println("tf.ReadPkg errror = ",err)
		return
	}

	//将mes的data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		
		//遍历当前在线用户列表
		for _,v := range loginResMes.UserIds{
			if v == userId {
				continue
			}
			fmt.Println("在线用户id:\t",v)
			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()

		//启动一个协程，接收服务器的推送消息
		go serverProcessMes(conn)

		

		//登陆成功，循环显示菜单
		for{
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}

	return 
}