package main

import (
	"fmt"
	"reflect"
)

func main() {
	// t := reflect.TypeOf(i)  //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	// v := reflect.ValueOf(i) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// 反射的字段可读写
	// 如果下面这样写，那么会发生错误
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // no new variables on left side of :=

	// 如果要修改相应的值，必须传递指针。
	var y float64 = 3.4
	p := reflect.ValueOf(&y)
	p := p.Elem()
	p.SetFloat(7.1)

}
