package main
import(
	"fmt"
	"os"
	"go_code/hello_project/chapter16chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main(){

	//用户选择
	var key int
	//是否继续显示菜单

	for true {
		fmt.Println("-------------------欢迎登陆聊天室系统--------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册")
		fmt.Println("\t\t\t 3 退出")
		fmt.Println("\t\t\t 4 请选择（1-3）")

		fmt.Scanln(&key)
		switch key{
			case 1 : 
				fmt.Println("登陆聊天室")
				fmt.Println("请输入用户id")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码")
				fmt.Scanln(&userPwd)
				up := &process.UserProcessor{}
				up.Login(userId,userPwd)
			case 2 :
				fmt.Println("注册用户")
				fmt.Println("请输入用户id:")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码：")
				fmt.Scanln(&userPwd)
				fmt.Println("请输入用户名称：")
				fmt.Scanln(&userName)
				up := &process.UserProcessor{}
				up.Register(userId,userPwd,userName)
				
			case 3 :
				fmt.Println("退出系统")
				//终止程序
				os.Exit(0)
			default :
				fmt.Println("您的输入有误，请重新输入")
		}
	}


}