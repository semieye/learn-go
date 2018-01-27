package main

import "fmt"

// 返回多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 闭包 匿名函数
// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 方法定义
// func (variable_name variable_data_type) function_name() [return_type]{
//    /* 函数体*/
// }

// Circle 定义结构体
type Circle struct {
	radius float64
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

// PI π
const PI = 3.14159265

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	// c.radius 即为 Circle 类型对象中的属性
	return PI * c.radius * c.radius
}
