
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



// res, err := Db.Exec("INSERT INTO userinfo (username, password, department,email) VALUES (?, ?, ?,?)","wd","123","it","wd@163.com")
// _, err1 := Db.Exec("update userinfo set username = ? where uid = ?","jack",1)
// _, err2 := Db.Exec("delete from userinfo where uid = ? ", 1)

// rows, err := Db.Query("SELECT username,password,email FROM userinfo")
// for rows.Next() { 
// 	var username,password,email string
// 	err = rows.Scan(&username, &password, &email)
// }

// row := Db.QueryRow("SELECT username,password,email FROM userinfo where uid = ?",1)
// var username,password,email string
// err :=row.Scan(&username,&password,&email)

// var stus []stu
// err := Db.Select(&stus,"SELECT username,password,email FROM userinfo")

// var s stu
// err1 := Db.Get(&s,"SELECT username,password,email FROM userinfo where uid = ?",2)