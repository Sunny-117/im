package message 

const(
	LoginMesType 				= "LoginMes"
	LoginResMesType 			= "LoginResMes"
	RegisterMesType 			= "RegisterMes"
	RegisterResMesType 			= "RegisterResMes"
	NotifyUserStatusMesType 	= "NotifyUserStatusMes"
	SmsMesType					= "SmsMesType"
)

//定义用户状态常量
const(
	UserOnline = iota
	UserOffline 
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`
}

//Data类型
type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `userName:"userName "`
 }

 type LoginResMes struct{
	 Code int `json:"code"`
	 UserIds []int  //保存当前在线用户的id
	 Error string `json:"error"`
 }

 type RegisterMes struct{
	 User User `json:"code"`
 }

 type RegisterResMes struct{
	 Code int 		`json:"code"` //返回状态码
	 Error string 	`json:"error"`
 }

 //推送消息
 type NotifyUserStatusMes struct{
	UserId int			`json:"userId"`
	Status int			`json:"status"`
 }

 //连通消息结构体
 type SmsMes struct{
	 Content string 			`json:"content"`
	 User    /*继承关系*/ 		
 }