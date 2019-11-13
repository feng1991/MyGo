package controller

import(
	"fmt"
	"net/http"

	"simpleWeb/app/common"
)

type TestController struct{
	W http.ResponseWriter
	R *http.Request
}

func (c *TestController) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in TestController\n")
	title := common.ConvertToString("zhongguo", "gbk", "utf-8")
	fmt.Fprintf(c.W,title)
}