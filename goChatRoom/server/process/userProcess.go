package process2
import(
	"fmt"
	"net"
	"go_code/hello_project/chapter16chatroom/common/message"
	"go_code/hello_project/chapter16chatroom/server/utils"
	"go_code/hello_project/chapter16chatroom/server/model"
	"encoding/json"
)

type UserProcess struct{
	Conn net.Conn
	UserId int //用户id来标识连接
}

//通知通知在线的用户的方法，sserId上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int){
	//遍历在线用户的列表
	for id,up := range MyUserMgr.onlineUsers {
		//过滤到自己
		if id == this.UserId {
			continue
		}
		//通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int){
	//构建通知消息的报文
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data , err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyStatus json.Marshal error = ",err)
		return
	}
	mes.Data = string(data)

	data , err = json.Marshal(mes)
	if err != nil {
		fmt.Println("notifyStatus json.Marshal error = ",err)
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("notifyMeOnline error = ",err)
		return
	}

}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error){
	var loginMes message.LoginMes
	//反序列化报文的data
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Ummarshal fail error = ",err)
		return 
	}

	var resMes message.Message
	resMes.Type = message.LoginMesType
	
	//登录返回对象
	var loginResMes message.LoginResMes

	//模拟查询数据库
	//如果用户id = 100 ， 密码 = 123456
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
		
	// }else{
	// 	//不合法
	// 	loginResMes.Code = 500 //表示该用户不存在
	// 	loginResMes.Error = "该用户不存在，请注册再使用..."
	// }
	user , err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	
	if err != nil{

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 //表示该用户不存在
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403 //表示该用户不存在
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code = 505 //表示该用户不存在
			loginResMes.Error = "服务器内部错误..."
		}
	}else{
		loginResMes.Code = 200
		//把用户放入onlinesUsers中
		this.UserId = loginMes.UserId
		
		MyUserMgr.AddOnlineUser(this)
		//通知其他在线用户：我上线了
		this.NotifyOthersOnlineUser(this.UserId)
		for id , _ := range MyUserMgr.onlineUsers{
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}
		fmt.Println(user,"登录成功")
	}

	data , err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json Marshal error = " , err)
		return
	}
	resMes.Data = string(data)
	data , err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json Marshal error = " , err)
		return
	}
	//6.发送数据包
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}


func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error){
	var registerMes message.RegisterMes
	//反序列化成结构体,message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes


	err = model.MyUserDao.Register(&registerMes.User)
	//说明找到用户已存在
	if err != nil{
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}

	//返回消息给客户端
	data,err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("josn.Marshal fail , error = ",err)
		return
	}
	resMes.Data = string(data)
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("josn.Marshal fail , error = ",err)
		return
	}

	//发送数据
	tf := &utils.Transfer{
		Conn :this.Conn,
	}
	err = tf.WritePkg(data)
	return


}
