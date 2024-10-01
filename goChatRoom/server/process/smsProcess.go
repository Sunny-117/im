package process2
import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"go_code/hello_project/chapter16chatroom/server/utils"
	"net"
	"encoding/json"
)

type SmsProcess struct{
}

func (this *SmsProcess) SendGroupMes(mes *message.Message){

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal error , ", err)
		return
	}

	data , err := json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal error = ", err)
	}

	//遍历在线用户，转发消息
	for id,up := range MyUserMgr.onlineUsers {
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToOtherOnlineUsers(data,up.Conn)
	}
}


func (this *SmsProcess) SendMesToOtherOnlineUsers(data []byte,conn net.Conn) {

	tf := &utils.Transfer{
		Conn :conn,
	}
	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("转发消息失败...，, error = ",err)
		return
	}
}