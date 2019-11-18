package lib
 
import (
	"fmt"
	"net/http"
	"reflect"
	
	"simpleWeb/config"
	"simpleWeb/router"
)

type App struct{
	W http.ResponseWriter
	R *http.Request
}

func init() {
}


//启动服务端
func (app *App) Setup() {
	//设置相应函数
	http.HandleFunc("/", app.IndexHandler)
	//开启服务监听请求
	fmt.Println("Start http server at " + config.SERVER_IP_PORT)
    http.ListenAndServe(config.SERVER_IP_PORT, nil)
}


//响应请求
func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	app.W = w
	app.R = r
	// 处理请求
	query := r.URL.Query()
	c := query.Get("c")
	a := query.Get("a")
	// 利用反射给struct赋值和调用方法
	route := router.RegStruct[c]
	if route == nil {
		fmt.Fprintf(w,"No such controller!")
		return
	}
	ele := reflect.ValueOf(route)
	ele.Elem().FieldByName("W").Set(reflect.ValueOf(w))
	ele.Elem().FieldByName("R").Set(reflect.ValueOf(r))  
	method := ele.MethodByName(a)
	if !method.IsValid() {
	    fmt.Fprintf(w,"No such method!")
		return
	}
	//如果传参，需用reflect.ValueOf()处理
	method.Call([]reflect.Value{})
}