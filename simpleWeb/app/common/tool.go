package common

import (
	"reflect"

	"github.com/axgle/mahonia"
)


// 编码转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 动态调用方法
func DynamicInvoke(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	//动态调用方法
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	//动态访问属性
	//reflect.ValueOf(object).Elem().FieldByName("Name")
}
