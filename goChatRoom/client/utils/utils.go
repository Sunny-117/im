package utils
import(
	"fmt"
	"net"
	"go_code/hello_project/chapter16chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

type Transfer struct{
	Conn net.Conn
	Buf [8096]byte //传输时使用的缓存
}

func (this *Transfer) WritePkg(data []byte) (err error){
	//先发送数据包的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	// buf = make([]byte,4)
	binary.BigEndian.PutUint32(this.Buf[:4],pkgLen)
	//conn只能发送切片，先发送消息的长度
	n , err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.write len error",err)
		return
	}

	//发送data本身
	n , err = this.Conn.Write(data)
	if n!= int(pkgLen) || err != nil{
		fmt.Println("conn.write data error = ",err)
		return
	}
	return

}

func (this *Transfer) ReadPkg() (mes message.Message , err error){
	//buf := make([]byte,8089)
	//读取对端数据
	//先读前4个字节
	_ ,err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	// fmt.Println("数据包长度为：",buf[:4])
	//解析出报文长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	
	n,err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn read error = ",err)
		return
	}

	//反序列化
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal error")
		return
	}

	return
	
}