package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	// 与C不同之处：
	// 首先，像 %d 这样的数值格式并不接受表示符号或大小的标记， 打印例程会根据实参的类型来决定这些属性。
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	// 若你只想要默认的转换，如使用十进制的整数，你可以使用通用的格式 %v（对应“值”）；
	// 其结果与 Print 和 Println 的输出完全相同。此外，这种格式还能打印任意值，甚至包括数组、结构体和映射。
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Printf("%v\n", timeZone) // 或只用 fmt.Println(timeZone)

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	// 请注意其中的&符号）当遇到 string 或 []byte 值时， 可使用 %q 产生带引号的字符串；
	// 而格式 %#q 会尽可能使用反引号。 （%q 格式也可用于整数和符文，它会产生一个带单引号的符文常量。）
	// 此外，%x 还可用于字符串、字节数组以及整数，并生成一个很长的十六进制字符串，
	// 而带空格的格式（% x）还会在字节之间插入空格。

	// 另一种实用的格式是 %T，它会打印某个值的类型.
	fmt.Printf("%T\n", timeZone)

	// 若你想控制自定义类型的默认格式，只需为该类型定义一个具有 String() string 签名的方法。
	// 对于我们简单的类型 T，可进行如下操作。
	fmt.Printf("%v\n", t)
	// 会打印出如下格式：
	// 7/-2.35/"abc\tdef"
}

// T test struct
type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
