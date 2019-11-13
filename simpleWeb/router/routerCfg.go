package router

import(
	"simpleWeb/app/controller"
)

func init() {
	//注册每个要用到的控制器
	addRouter("Index",&controller.IndexController{})
	addRouter("Test",&controller.TestController{})
}
