package controller

import(
	"fmt"
	"net/http"
)

type IndexController struct{
	W http.ResponseWriter
	R *http.Request
}

func (c *IndexController) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in IndexController\n")
}