package main

import "fmt"

/*
for init; condition; post { } // 如同C的for循环
for condition { }             // 如同C的while循环
for { }                       // 如同C的for(;;)循环
*/

func main() {
	//if 语句后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
	/* 局部变量定义 */
	a := 100

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
	// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break。
	/* 定义局部变量 */
	var grade = "B"
	var marks = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T\n", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	// select语句
	// select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
	// select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
	// 每个case都必须是一个通信
	// 所有channel表达式都会被求值
	// 所有被发送的表达式都会被求值
	// 如果任意某个通信可以进行，它就执行；其他被忽略。
	// 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	// 否则：
	// 如果有default子句，则执行该语句。
	// 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// Go没有逗号操作符，而 ++ 和 -- 为语句而非表达式。 因此，若你想要在 for 中使用多个变量，
	// 应采用平行赋值的方式 （因为它会拒绝 ++ 和 --）
	w := [5]int{0, 1, 2, 3, 4}
	// 反转 w
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	fmt.Println(w)

	// 解析UTF-8， 将每个独立的Unicode码点分离出来。错误的编码将占用一个字节，range自动以符文U+FFFD来代替。
	for pos, char := range "日本\x80語" { // \x80 是个非法的UTF-8编码
		fmt.Printf("字符 %#U 始于字节位置 %d\n", char, pos)
	}
}

// 若 switch 后面没有表达式，它将匹配 true，因此，我们可以将 if-else-if-else 链写成一个 switch
// break + 标签 可以跳出switch和多重循环
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// switch 并不会自动下溯，但 case 可通过逗号分隔来列举相同的处理条件。
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

// Compare 按字典顺序比较两个字节切片并返回一个整数。
// 若 a == b，则结果为零；若 a < b；则结果为 -1；若 a > b，则结果为 +1。
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
