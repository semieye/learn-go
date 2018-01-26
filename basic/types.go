//my first go program.
package main

import (
	"fmt"
)

//示例代码
var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func main() {
	//定义一个名称为“variableName”，类型为"type"的变量
	var a int64 = 1
	//定义三个类型都是“type”的变量
	// var vname1, vname2, vname3 type

	/*
		定义三个变量，它们分别初始化为相应的值
		vname1为v1，vname2为v2，vname3为v3
		然后Go会根据其相应值的类型来帮你初始化它们
	*/
	var b, c, d = 2, "test", 0.45

	/*
		定义三个变量，它们分别初始化为相应的值
		vname1为v1，vname2为v2，vname3为v3
		编译器会根据初始化的值自动推导出相应的类型
	*/
	e, f, g := 1e10, "中国", "[1, 2]"

	// _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：
	_, h := 34, 35

	const constantName = "value"
	//如果需要，也可以明确指定常量的类型：
	const Pi float32 = 3.1415926
	// const Pi = 3.1415926
	const i = 10000
	const MaxThread = 10
	const prefix = "prefix_"

	// Go对于已声明但未使用的变量会在编译阶段报错
	fmt.Println(a, b, c, d, e, f, g, h)

	// call func test()
	test()
}

// for var define test
func test() {
	// 在Go中，布尔值的类型为bool，值是true或false，默认为false。

	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作

	// Go对于已声明但未使用的变量会在编译阶段报错
	fmt.Println(valid, available)

	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v", c)
}
