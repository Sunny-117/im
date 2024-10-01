package main
import(
	"fmt"
	"net"
	"go_code/hello_project/chapter16chatroom/common/message"
	"go_code/hello_project/chapter16chatroom/server/utils"
	"go_code/hello_project/chapter16chatroom/server/process"
	"io"
)

type Processor struct{
	Conn net.Conn
}


//负责分发数据包给具体函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error){
	switch mes.Type{
		case message.LoginMesType : 
			//处理登录
			//创建UserProcess实例
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType :
			//处理注册
			fmt.Println("处理注册请求...")
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType :
			//转发群聊消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
			fmt.Println("mes = ",mes)
		default :
			fmt.Println("消息类型不存在，无法处理...")
	}
	return 
}

func (this *Processor) process2() (err error){
	
	//创建实例
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	for{
		//读取报文
		mes , err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			}else{
				// err = errors.New("read pkg header error")
				fmt.Println("readPkg error = ",err)
				return err
			}
			
		}
		// fmt.Println(mes)
		//处理数据包
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}

	return
}