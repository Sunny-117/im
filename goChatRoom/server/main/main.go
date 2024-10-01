package main
import(
	"fmt"
	"net"
	"time"
	"go_code/hello_project/chapter16chatroom/server/model"
)

//处理连接
func process(conn net.Conn){
	defer conn.Close()

	//processor结构体代表一个连接类
	processor := &Processor{
		Conn : conn,
	}

	//处理通讯
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端协议错误，error = ",err)
		return
	}
}


//对userDao初始化
func initUserDao(){
	//这个pool就是redis.go文件里定义的全局变量
	model.MyUserDao = model.NewUserDao(pool)
}


func main(){

	initPool("localhost:6379",16,0,300*time.Second)
	initUserDao()
	//提示信息
	fmt.Println("服务器[分层]监听8889端口...")
	listen , err := net.Listen("tcp","localhost:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	}

	//监听连接
	for{
		fmt.Println("等待连接...")
		conn , err := listen.Accept()
		if err != nil{
			fmt.Println("listen.Accept error = ",err)
		}

		//用协程去与客户端通讯
		go process(conn)
	}
}