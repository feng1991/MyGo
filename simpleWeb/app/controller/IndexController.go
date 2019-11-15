package controller

import(
	"fmt"

	"simpleWeb/lib/base"
)

type IndexController struct{
	base.Controller
}

func (c *IndexController) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in IndexController\n")
}