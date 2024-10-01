package main
import(
	"github.com/garyburd/redigo/redis"
	"time"
)

//全局连接池
var pool *redis.Pool

func initPool(address string, maxIdle , maxActive int , idleTimeout time.Duration){
	pool = &redis.Pool{
		MaxIdle : maxIdle,
		MaxActive : maxActive,
		IdleTimeout : idleTimeout,
		Dial : func()(redis.Conn , error){
			return redis.Dial("tcp",address)
		},
	}
}
