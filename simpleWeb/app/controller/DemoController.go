package controller

import(
	"fmt"

	"simpleWeb/app/common"
	"simpleWeb/app/model"
	"simpleWeb/lib/base"
)

type DemoController struct{
	base.Controller
}

type Log struct{
	Id int `json:"-"` //不参与转json
	Name string
	Value string
}

func (c *DemoController) GetName(){
	fmt.Fprintf(c.W,"\nHere is GetName in DemoController\n")
	title := common.ConvertToString("zhongguo", "gbk", "utf-8")
	fmt.Fprintf(c.W,title)
}

func (c *DemoController) GetDb(){
	var db = new(model.CityModel)
	db.InitModel()
	city,err := db.GetCity()
	if err != nil{
		fmt.Fprintln(c.W,err)
		return
	}
	fmt.Fprintln(c.W,city)
}

func (c *DemoController) GetJson(){
	//obj := Log{Id:1,Name:"jack",Value:"a teacher"}
	obj := make(map[string]interface{})
	obj["name"] = "jack"
	obj["sex"] = "male"
	fmt.Fprintf(c.W,"Type:%T \n",obj)
	data := c.OutputJson(obj)
	fmt.Fprintln(c.W,data)
}

func (c *DemoController) GetParam(){
	obj := c.Param("name")
	fmt.Fprintf(c.W,"get:%v \n",obj)
}

func (c *DemoController) GetHead(){
	arr := c.GetHeader("User-Agent")
	for _,v := range arr{
		fmt.Fprintf(c.W,v)
	}
}