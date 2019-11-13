package router
 
import (

)

//用于保存实例化的结构体对象
var RegStruct = make(map[string]interface{})

func addRouter(name string,ctl interface{}) {
	RegStruct[name] = ctl
}
