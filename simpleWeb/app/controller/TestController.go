package controller

import(
	"fmt"
	"net/http"

	"simpleWeb/app/common"
	"simpleWeb/app/model"
	"simpleWeb/lib/base"
)

type TestController struct{
	W http.ResponseWriter
	R *http.Request
	base.Controller
}

func (c *TestController) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in TestController\n")
	title := common.ConvertToString("zhongguo", "gbk", "utf-8")
	fmt.Fprintf(c.W,title)
}

func (c *TestController) GetDb(){
	var db = new(model.CityModel)
	db.InitModel()
	city,err := db.GetCity()
	if err != nil{
		fmt.Fprintln(c.W,err)
		return
	}
	fmt.Fprintln(c.W,city)
}