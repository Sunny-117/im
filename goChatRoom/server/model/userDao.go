package model

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"go_code/hello_project/chapter16chatroom/common/message"

)

var(
	MyUserDao *UserDao
)

//定义UserDao结构体
type UserDao struct{
	pool *redis.Pool
}

//工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	//因为返回的是指针，所以要分配内存？函数的返回列表内不会为指针分配地址
	userDao = &UserDao{
		pool : pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn , id int) (user *User,err error){

	//通过id查询redis
	res, err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		//错误
		if err == redis.ErrNil{ //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{} //这一步是多余的？
	//反实例化User实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("json Unmarshal error = ",err)
		return
	}
	return

}

//完成登录的验证
//如果登录成功，则返回user实例
func (this *UserDao) Login(userId int , userPwd string) (user *User,err error){

	conn := this.pool.Get()
	defer conn.Close()
	user , err = this.getUserById(conn,userId)
	if err != nil{
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}


func (this *UserDao) Register(user *message.User) (err error){
	//获取Redis连接
	conn := this.pool.Get()
	defer conn.Close()

	_ , err = this.getUserById(conn,user.UserId)
	if err == nil{
		//没有错误，说明用户已存在
		err = ERROR_USER_EXISTS 
		return
	}

	//说明用户id在Redis还没有，可以完成注册
	data , err := json.Marshal(user)
	if err != nil {
		return
	} 

	//入库
	_ , err = conn.Do("HSet","users",user.UserId,string(data))
	if err != nil{
		fmt.Println("往reids写入用户错误，error = ",err)
		return
	}
	return
}