// go语言 基础类型
package main

import (
	"errors"
	"fmt"
)

//go keywords
// break    default      func    interface    select
// case     defer        go      map          struct
// chan     else         goto    package      switch
// const    fallthrough  if      range        type
// continue for          import  return       var

// 示例代码
var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func main() {
	// 变量定义
	// var vname1, vname2, vname3 type
	var a int64 = 1
	// Go会根据其相应值的类型来初始化它们
	var b, c, d = 2, "test", 0.45
	e, f, g := 1e10, "中国", "China"

	// _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。

	_, h := 34, 35

	// 常量定义
	const constantName = "value"
	// 如果需要，也可以明确指定常量的类型：
	const Pi float32 = 3.1415926 // const Pi = 3.1415926
	const MaxThread = 10
	const prefix = "prefix_"

	// Go对于已声明但未使用的变量会在编译阶段报错
	fmt.Println(a, b, c, d, e, f, g, h)

	// 在Go中，布尔值的类型为bool，值是true或false，默认为false。
	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作
	fmt.Println(valid, available)
	// 复数类型
	var complex complex64 = 5 + 5i
	fmt.Printf("complex64 is: %v\n", complex)

	// 字符串
	var frenchHello string                 // 声明变量为字符串的一般方法
	var emptyString string = ""            // 声明了一个字符串变量，初始化为空字符串
	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"           // 同上
	frenchHello = "Bonjour"                // 常规赋值
	emptyString = "N/A"
	fmt.Println(frenchHello, emptyString, no, yes, maybe, japaneseHello)

	// 修改字符串
	s := "hello"
	bv := []byte(s) // 将字符串 s 转换为 []byte 类型
	bv[0] = 'C'
	fmt.Printf("%s\n", string(bv))
	// 字符串连接
	fmt.Printf("%s\n", "hello,"+" world")
	// 字符串切片
	fmt.Printf("%s\n", s[1:])
	// 多行字符串
	fmt.Printf("%s\n", `hello
	                          world`)

	err := errors.New("panic: stackoverflow")
	if err != nil {
		fmt.Print(err)
	}

	// 变量声明分组
	const (
	// i = 100
	// pi = 3.1415
	// prefix = "Go_"
	)

	var (
	// i      int
	// pi     float32
	// suffix string
	)

	// itoa 常量计数器
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 常量声明省略值时，默认和之前一个值的字面相同。
		// 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)
	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
	const (
		hh, ii, jj = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	)
	const (
		aa         = iota //a=0
		bb         = "B"
		cc         = iota             //c=2
		dd, ee, ff = iota, iota, iota //d=3,e=3,f=3
		gg         = iota             //g = 4
	)
	// 除非被显式设置为其它值或iota，每个const分组的第一个常量被默认设置为它的0值，
	// 第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，则它也被设置为iota。
	fmt.Printf("\n")
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, x, y, z, w, v)

	fmt.Printf("大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；\n")
	fmt.Printf("小写字母开头的就是不可导出的，是私有变量。\n")
	fmt.Printf("大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；\n")
	fmt.Printf("小写字母开头的就是有private关键词的私有函数。\n")
}
