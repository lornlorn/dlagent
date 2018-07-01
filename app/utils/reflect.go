package utils

import (
	"app/api"
	"errors"
	"fmt"
	"log"
	"reflect"
)

// FunctionMap 定义函数映射类型
type FunctionMap map[string]reflect.Value

// FuncMap 声明函数映射
var FuncMap FunctionMap

/*
InitFunctionMap func()
初始化函数映射表
*/
func InitFunctionMap() {
	var ajaxapi api.API
	FuncMap = make(FunctionMap, 0)
	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	apiObj := reflect.ValueOf(&ajaxapi)
	apiObjType := apiObj.Type()
	//读取方法数量
	methodNum := apiObj.NumMethod()
	fmt.Println("NumMethod:", methodNum)

	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < methodNum; i++ {
		methodName := apiObjType.Method(i).Name
		fmt.Println("index:", i, " MethodName:", methodName)
		FuncMap[methodName] = apiObj.Method(i) //<<<
	}
}

// FuncCall func(mName string, param ...interface{})
func FuncCall(mName string, param ...interface{}) []reflect.Value {

	params := make([]reflect.Value, len(param))

	for i, v := range param {
		log.Printf("index : %v, %v\n", i, v)
		params[i] = reflect.ValueOf(v)
	}

	//使用方法名字符串调用指定方法
	return FuncMap[mName].Call(params)

}

// ReflectCall func(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error)
func ReflectCall(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}
