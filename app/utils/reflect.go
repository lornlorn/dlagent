package utils

import (
	"app/api"
	"errors"
	"log"
	"reflect"
)

// FunctionMap 定义函数映射类型
type FunctionMap map[string]reflect.Value

// type FunctionMap map[string]interface{}

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
	// fmt.Println("NumMethod:", methodNum)

	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < methodNum; i++ {
		methodName := apiObjType.Method(i).Name
		log.Printf("Index : %v, MethodName : %v\n", i, methodName)
		FuncMap[methodName] = apiObj.Method(i)
	}
}

/*
FuncCall func(mName string, param ...string)
param指定string类型,若有其他类型需求，调用方法内自行转换
*/
func FuncCall(mName string, params ...interface{}) []reflect.Value {

	f := reflect.ValueOf(FuncMap[mName])
	if len(params) != f.Type().NumIn() {
		log.Println("The number of input params not match")
		return nil
	}

	in := make([]reflect.Value, len(params))

	for k, v := range params {
		log.Printf("index : %v, %v\n", k, v)
		in[k] = reflect.ValueOf(v)
	}

	//使用方法名字符串调用指定方法
	return f.Call(in)

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
