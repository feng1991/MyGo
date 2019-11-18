
package base

import(
	//"fmt"
	"encoding/json"
	"net/http"

	_ "github.com/axgle/mahonia"
)

type Controller struct{
	W http.ResponseWriter
	R *http.Request
}

// 获得get或post的传参
func (ctl *Controller) Param(param string) string {
	p := ctl.R.FormValue(param)
	return p
}

// 输出json格式
func (ctl *Controller) OutputJson(obj interface{}) string {
	//data, err := json.MarshalIndent(obj,"","   ")
	data, err := json.Marshal(obj)
	if err != nil {
		panic("JSON marshaling failed: "+err.Error())
		return ""
	}
	return string(data)
}

// 获得头部信息
func (ctl *Controller) GetHeader(name string) []string {
	//var str string
	// for i,v := range ctl.R.Header {
	// 	str += fmt.Sprintf("%v : %v\n",i,v)
	// }
	//return str
	return ctl.R.Header[name]
}
