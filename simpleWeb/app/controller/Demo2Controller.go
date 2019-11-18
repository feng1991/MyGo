package controller

import(
	"fmt"

	"simpleWeb/lib/base"
)

type Demo2Controller struct{
	base.Controller
}

func (c *Demo2Controller) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in Demo2Controller\n")
}