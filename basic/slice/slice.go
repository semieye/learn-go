// go语言 动态数组
package main

import (
	"fmt"
)

func main() {
	// slice并不是真正意义上的动态数组，而是一个引用类型。
	// slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
	// 和声明array一样，只是少了长度
	var fslice []byte
	slice := []byte{'a', 'b', 'c', 'd'}
	printSliceByte(fslice)
	printSliceByte(slice)
	// slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，
	// 其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

	// ** i 应该是slice指向的元素下标，j应该是取值取到的下标（不包括此下标元素）**

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	// var aa, bb []byte
	aa := ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	bb := ar[3:5]
	// b的元素是：ar[3]和ar[4]
	printSliceByte(aa)
	printSliceByte(bb)

	// slice有一些简便的操作
	// slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	// slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	// 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，
	// 第二个是数组的长度，即等价于ar[0:len(ar)]

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte
	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素
	// 从slice中获取slice
	aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4 指向下标3，取出下标7之前的
	printSliceByte(aSlice)
	// cap() == 10 - 3 = 7 aSlice指向array之后，cap就是第二个序列？
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	printSliceByte(bSlice)
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	printSliceByte(bSlice)
	bSlice = aSlice[:] // bSlice包含所有aSlice的元素: d,e,f,g

	// slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，
	// 例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。

	// 对于slice有几个有用的内置函数：
	// len 获取slice的长度
	// cap 获取slice的最大容量
	// append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	// copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

	// 注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	// 但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
	// 返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。

	// 从Go1.2开始slice支持了三个参数的slice，
	// 之前我们一直采用这种方式在slice或者array基础上来获取一个slice
	// var array [10]int
	// slice := array[2:4]
	// 这个例子里面slice的容量是8，新版本里面可以指定这个容量
	// slice = array[2:4:7]
	// 上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。
	// 如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSliceInt(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSliceInt(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSliceInt(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSliceInt(number3)

	var numbers2 []int
	printSliceInt(numbers)

	/* 允许追加空切片 */
	numbers2 = append(numbers2, 0)
	printSliceInt(numbers2)

	/* 向切片添加一个元素 */
	numbers2 = append(numbers2, 1)
	printSliceInt(numbers2)

	/* 同时添加多个元素 */
	numbers2 = append(numbers2, 2, 3, 4)
	printSliceInt(numbers2)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers3 := make([]int, len(numbers2), (cap(numbers2))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers3, numbers2)
	printSliceInt(numbers3)
}

func printSliceByte(x []byte) {
	//这里的len()是用来获得现在切片数据的长度，而cap()则为了求总容量
	fmt.Printf("len=%d cap=%d slice=%c\n", len(x), cap(x), x)
}

func printSliceInt(x []int) {
	//这里的len()是用来获得现在切片数据的长度，而cap()则为了求总容量
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
