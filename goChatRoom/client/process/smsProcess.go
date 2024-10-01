package process

import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/json"
	"go_code/hello_project/chapter16chatroom/client/utils"
)

type SmsProcess struct{
}

func (this *SmsProcess) SendGroupMes(content string) (err error){

	//创建mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserStatus = curUser.UserStatus
	smsMes.UserId = curUser.UserId

	//序列化
	data ,err := json.Marshal(smsMes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal error = " , err.Error())
		return
	}
	mes.Data = string(data)

	data ,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal error = " , err.Error())
		return
	}
	tf := &utils.Transfer{
		Conn : curUser.Conn,
	}

	//发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg error = ",err.Error())
		return
	}
	return
	
}