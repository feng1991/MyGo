
package cache

import(
	"fmt"

	"simpleWeb/config"

	"github.com/gomodule/redigo/redis"
)

type SimpleRedis struct{
	Conn redis.Conn
}

// 连接redis
func (re *SimpleRedis) InitRedis(){
	dialStr := fmt.Sprintf("%s:%d",config.REDIS_SERVER,config.REDIS_PORT)
	c,err := redis.Dial("tcp",dialStr)
    if err != nil {
        panic("connect redis error :"+err.Error())
        return
	}
	re.Conn = c
	//defer re.conn.Close()
	_,err = re.Conn.Do("AUTH", config.REDIS_AUTH)
	if err != nil {
        panic("AUTH redis error :"+err.Error())
        return
	}
}


// 关闭redis
func (re *SimpleRedis) Close(){
	re.Conn.Close()
}


// 字符串存储
func (re *SimpleRedis) Set(name string, value string) error {
	_, err := re.Conn.Do("SET", name, value)
	return err
}


// 带过期时间的字符串存储
func (re *SimpleRedis) SetWithTtl(name string, value string, second int) error {
	_, err := re.Conn.Do("SET", name, value, "EX", second)
	return err
}


// 字符串读取
func (re *SimpleRedis) Get(name string) (string,error) {
	value, err := redis.String(re.Conn.Do("GET", name))
    return value,err
}


// 过期时间设置
func (re *SimpleRedis) Expire(name string, second int) error {
	_, err := re.Conn.Do("expire", name, second)
	return err
}


// 批量字符串存取
//_, err := re.Conn.Do("MSET", "name", "wd","age",22)
//res, err := redis.Strings(re.Conn.Do("MGET", "name","age"))

// 队列存取
//_, err = re.Conn.Do("LPUSH", "list1", "ele1","ele2","ele3")
//res, err := redis.String(re.Conn.Do("LPOP", "list1"))

// 哈希存取
//_, err = re.Conn.Do("HSET", "student","name", "wd","age",22)
//res, err := redis.Int64(re.Conn.Do("HGET", "student","age"))