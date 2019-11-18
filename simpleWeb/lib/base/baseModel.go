
package base

import(
	"fmt"
	"time"
	"database/sql"
	
	"simpleWeb/config"

	_ "github.com/go-sql-driver/mysql"
)

type Model struct{
	Conn *sql.DB
}

// 连接db
func (model *Model) InitModel() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",config.USERNAME,config.PASSWORD,config.NETWORK,config.SERVER,config.PORT,config.DATABASE)
	db,err := sql.Open("mysql",dsn)
	if err != nil{
		panic("Open mysql failed,err:"+err.Error())
		return
	}
	model.Conn = db
	model.Conn.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	model.Conn.SetMaxOpenConns(100)//设置最大连接数
	model.Conn.SetMaxIdleConns(16) //设置闲置连接数
}

