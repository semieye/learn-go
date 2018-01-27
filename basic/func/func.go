package main

import (
	"fmt"
	"io"
	"os"
)

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

	// defer test
	bar()
}

// PI π
const PI = 3.14159265

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	// c.radius 即为 Circle 类型对象中的属性
	return PI * c.radius * c.radius
}

// 可命名结果形参
/*
func ReadFull(r Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
*/

// Go的 defer 语句用于预设一个函数调用（即推迟执行函数），该函数会在执行 defer 的函数返回之前立即执行。
// 它显得非比寻常， 但却是处理一些事情的有效方式，例如无论以何种路径返回，都必须释放资源的函数。
// 典型的例子就是解锁互斥和关闭文件。

// Contents 将文件的内容作为字符串返回。
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close 会在我们结束后运行。

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append 将在后面讨论。
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // 我们在这里返回后，f 就会被关闭。
		}
	}
	return string(result), nil // 我们在这里返回后，f 就会被关闭。
}

/*
被推迟函数的实参（如果该函数为方法则还包括接收者）在推迟执行时就会求值， 而不是在调用执行时才求值。
这样不仅无需担心变量值在函数执行时被改变， 同时还意味着单个已推迟的调用可推迟多个函数的执行。
下面是个简单的例子。
for i := 0; i < 5; i++ {
	defer fmt.Printf("%d ", i)
}
被推迟的函数按照后进先出（LIFO）的顺序执行，因此以上代码在函数返回时会打印 4 3 2 1 0。

一个更具实际意义的例子是通过一种简单的方法， 用程序来跟踪函数的执行。我们可以编写一对简单的跟踪例程：

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// 像这样使用它们：
func a() {
	trace("a")
	defer untrace("a")
	// 做一些事情....
}
*/

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func foo() {
	defer untrace(trace("foo"))
	fmt.Println("in foo")
}

func bar() {
	defer untrace(trace("bar"))
	fmt.Println("in bar")
	foo()
}

// 数据
// Go提供了两种分配原语，即内建函数 new 和 make。
// new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。

// File the struct of file
type File struct {
	fd      int
	name    string
	dirinfo *int
	nepipe  int
}

// NewFile Test
// 若复合字面不包括任何字段，它将创建该类型的零值。表达式 new(File) 和 &File{} 是等价的。
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0} //复合字面 该表达式在每次求值时都会创建新的实例
	return &f
	// return &File{fd, name, nil, 0}   // 每当获取一个复合字面的地址时，都将为一个新的实例分配内存
	// return &File{fd: fd, name: name} // 未给出的字段值将赋予零值
}

// 复合字面 初始化 数组 切片 映射
// a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
// s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
// m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

// 内建函数 make(T, args) 的目的不同于 new(T)。它只用于创建切片、映射和信道，
// 并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值。
/*
下面的例子阐明了 new 和 make 之间的区别：

var p *[]int = new([]int)       // 分配切片结构；*p == nil；基本没用
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 习惯用法：
v := make([]int, 100)
请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。
*/

// Min ... 形参可指定具体的类型 其形参可为 ...int 类型
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // 最大的 int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
