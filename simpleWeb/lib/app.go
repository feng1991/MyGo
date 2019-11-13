package lib
 
import (
	"fmt"
	//"strings"
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
	app.getRequestInfo(w,r)
	// 处理请求
	query := r.URL.Query()
	c := query.Get("c")
	a := query.Get("a")
	// 利用反射给struct赋值和调用方法
	route := router.RegStruct[c]
	ele := reflect.ValueOf(route)
	ele.Elem().FieldByName("W").Set(reflect.ValueOf(w))
	ele.Elem().FieldByName("R").Set(reflect.ValueOf(r))  
	ele.MethodByName(a).Call([]reflect.Value{})
}



func (app *App) getRequestInfo(w http.ResponseWriter, r *http.Request) {
	// request信息
	// pathInfo := strings.Trim(r.URL.Path, "/")
	// parts := strings.Split(pathInfo, "/")
	// fmt.Fprintln(w,parts[0])
	//fmt.Fprintln(w,*r)
	fmt.Fprintf(w,"\nURL:%v,\nHeader:%v,\nBody:%v,\nHost:%v,\nPostForm:%v,\nRemoteAddr:%v\n\n",r.URL,r.Header,r.Body,r.Host,r.PostForm,r.RemoteAddr)
	// header信息
	for i,v := range r.Header {
		fmt.Fprintf(w,"%v : %v\n",i,v)
	}
}




// login := &loginController{}
// controller := reflect.ValueOf(login)
// method := controller.MethodByName(action)
// if !method.IsValid() {
//     method = controller.MethodByName(strings.Title("index") + "Action")
// }
// requestValue := reflect.ValueOf(r)
// responseValue := reflect.ValueOf(w)
// method.Call([]reflect.Value{responseValue, requestValue})


// func DynamicInvoke(object interface{}, methodName string, args ...interface{}) {
// 	inputs := make([]reflect.Value, len(args))
// 	for i, _ := range args {
// 		inputs[i] = reflect.ValueOf(args[i])
// 	}
// 	//动态调用方法
// 	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
// 	//动态访问属性
// 	reflect.ValueOf(object).Elem().FieldByName("Name")
// }