
package base

import(
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
