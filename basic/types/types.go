// 基础类型
// 包注释：会被godoc读取生成文档

package main

import (
	"fmt"
)

// 全局变量
var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

// 包中的main可以编译成可执行文件
func main() {
	// 变量定义 Go会根据其相应值的类型来初始化它们
	var a int64 = 1
	var b, c, d = 2, "test", 0.45
	e, f, g := 1e10, "飞机", "is flying"
	// _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。
	_, h := 34, 35

	// 常量定义 如果需要，也可以明确指定常量的类型
	const constantName = "value"
	const ( // 声明分组
		Pi     float32 = 3.1415926 // const Pi = 3.1415926
		prefix         = "prefix_"
	)

	// Go对于已声明但未使用的变量会在编译阶段报错
	fmt.Println(a, b, c, d, e, f, g, h)

	// 在Go中，布尔值的类型为bool，值是true或false，默认为false。
	var available bool // 一般声明
	valid := false     // 简短声明
	fmt.Println(valid, available)
	var complex complex64 = 5 + 5i // 复数类型
	fmt.Println(complex)

	// 字符串
	var frenchHello string                 // 声明变量为字符串的一般方法
	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"           // 同上
	frenchHello = "Bonjour"                // 常规赋值
	fmt.Println(frenchHello, no, yes, maybe, japaneseHello)

	// 修改字符串
	s := "hello"
	bv := []byte(s) // 将字符串 s 转换为 []byte 类型
	bv[0] = 'C'
	fmt.Printf("%s\n", string(bv))

	fmt.Printf("%s\n", "hello,"+" world")                        // 字符串连接
	fmt.Printf("%s\n", s[1:])                                    // 字符串切片
	fmt.Printf("%s\n", `hello
	                          world`) // 多行字符串
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
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, x, y, z, w, v)
	fmt.Println("大写字母开头的变量、函数是可导出的，小写字母开头的就是不可导出的。")
}
