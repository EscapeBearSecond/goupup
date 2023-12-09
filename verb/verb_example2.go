package main

import (
	"fmt"
	"reflect"
)

func main() {
	user := &User{
		Name:    "张三",
		Age:     18,
		Address: "北京",
		Phone:   "13012345678",
		Email:   "XXXXXXXXXXXXXXX",
		Gender:  "男",
	}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(obj interface{}) {
	// 获取type
	typ := reflect.TypeOf(obj)
	fmt.Printf("之前的type: %v\n", typ)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	fmt.Printf("之后的type: %v\n", typ)
	// 获取value
	val := reflect.ValueOf(obj)
	fmt.Printf("之前的value: %v\n", val)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	fmt.Printf("之后的value: %v\n", val)
	// 获取字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		fmt.Printf("JsonTag:%v", field.Tag.Get("json"))
	}
	// TODO:获取方法 这里有坑
	// for i := 0; i < typ.NumMethod(); i++ {
	// 	method := typ.Method(i)
	// 	fmt.Printf("%s: %v\n", method.Name, method.Type)
	// 	// 调用方法
	// 	method.Func.Call([]reflect.Value{val})
	// }
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
}

func (u *User) GetInfo() {
	fmt.Println(u)
}
