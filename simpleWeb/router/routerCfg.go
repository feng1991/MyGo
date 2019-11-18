package router

import(
	"simpleWeb/app/controller"
)

func init() {
	//注册每个要用到的控制器，访问链接如 http://IP:PORT/?c=Demo
	addRouter("Demo",&controller.DemoController{})
	addRouter("Demo2",&controller.Demo2Controller{})
}
