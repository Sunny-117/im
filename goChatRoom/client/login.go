package main
import(
	"fmt"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/json"
	_ "encoding/binary"
	"net"
	_ "time"
)

func login(userId int , userPwd string) (err error){

	//定协议...
	// fmt.Printf("userId = %d , userPwd = %s",userId,userPwd)

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
	

	err = writePkg(conn,data)
	if err != nil {
		fmt.Println("writePkg error = ", err)
		return
	}
	// //先发送data的长度
	// var pkgLen uint32
	// pkgLen = uint32(len(data))
	// var buf [4]byte
	// // buf = make([]byte,4)
	// binary.BigEndian.PutUint32(buf[:4],pkgLen)
	// //conn只能发送切片，先发送消息的长度
	// _ , err = conn.Write(buf[:4])
	// if err != nil{
	// 	fmt.Println("conn.write len error",err)
	// 	return
	// }

	// // fmt.Println("client send data len success, data len = " , len(data))

	// //发送消息本身
	// _ ,err = conn.Write(data)
	// if err != nil{
	// 	fmt.Println("conn.write data error",err)
	// 	return
	// }
	//休眠20秒
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20秒")
	mes , err = readPkg(conn)

	//将mes的data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("登录成功....")
	}else if loginResMes.Code == 500{
		fmt.Println(loginResMes.Error)
	}

	return 
}