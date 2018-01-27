// 数组类型
// 包注释：会被godoc读取生成文档

package main

import (
	"fmt"
)

func main() {
	// 数组定义： var arr [n]type
	// 在[n]type中，n表示数组的长度，type表示存储元素的类型。
	// 对数组的操作和其它语言类似，都是通过[]来进行读取或赋值。
	// 由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
	// 数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，
	// 而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。
	var arr [10]int                             // 声明了一个int类型的数组
	arr[0] = 42                                 // 数组下标是从0开始的
	arr[1] = 13                                 // 赋值操作
	fmt.Println("The first element is", arr[0]) // 获取数据，返回42
	fmt.Println("The last element is", arr[9])  //返回未赋值的最后一个元素，默认返回0

	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println("a b c:", a, b, c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("doubleArray, easyArray:", doubleArray, easyArray)
}
