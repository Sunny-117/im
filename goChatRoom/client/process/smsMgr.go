package process
import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/json"
)

func outputGroupMes(mes *message.Message){
	//反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Ummarshal error = ", err)
		return
	}

	//显示群发消息
	info := fmt.Sprintf("用户id:%d 对大家说：%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}