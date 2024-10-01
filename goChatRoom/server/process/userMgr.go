package process2
import (
	"fmt"
)

var (
	MyUserMgr *UserMgr
)

//维护在线用户
type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

func init(){
	MyUserMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

//对map进行增删改查
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DeleteOnlineUser(uid int) {
	delete(this.onlineUsers,uid)
}

//查询
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(uid int) (up *UserProcess,err error){
	//如何从map中取出一个值？
	up , ok := this.onlineUsers[uid]
	if !ok {
		//用户不在线 or 不存在
		err = fmt.Errorf("用户%d 不存在",uid)
		return
	}
	return
}