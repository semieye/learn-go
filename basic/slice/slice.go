// 切片：动态数组
// 包注释：会被godoc读取生成文档

package main

import (
	"fmt"
)

func main() {
	// slice并不是真正意义上的动态数组，而是一个引用类型。
	// slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
	// 和声明array一样，只是少了长度
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSliceInt("slice", slice)
	// slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，
	// ** i 应该是slice指向的元素下标，j应该是取值取到的下标（不包括此下标元素）**
	slice = slice[2:5]
	printSliceInt("slice[2:5]", slice)

	// slice有一些简便的操作
	// slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	// slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	// 如果从一个数组里面直接获取slice，可以这样ar[:]

	// slice内置函数：
	// len 获取slice的长度
	// cap 获取slice的容量
	// append 追加一个或者多个元素，然后返回同类型的slice
	// copy 右参数向左参数复制，返回复制的元素的个数

	// 演示一些简便操作
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var aSlice, bSlice []byte
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	printSliceByte("aSlice", aSlice)
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	printSliceByte("aSlice", aSlice)
	aSlice = array[:] // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素
	printSliceByte("aSlice", aSlice)
	aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4 指向下标3，取出下标7之前的
	printSliceByte("aSlice", aSlice)
	// 注意cap的值，后面注释有解释
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	printSliceByte("bSlice", bSlice)
	bSlice = aSlice[:3] // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	printSliceByte("bSlice", bSlice)
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	printSliceByte("bSlice", bSlice)
	bSlice = aSlice[:] // bSlice包含所有aSlice的元素: d,e,f,g
	printSliceByte("bSlice", bSlice)

	// slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，
	// 例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。

	// 注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	// 但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
	// 返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。

	// 从Go1.2开始slice支持了三个参数的slice，
	// slice = array[2:4:7] // 容量是7-2，即5。

	// 演示自动扩容，cap容量变化很奇怪的。
	var numbers []int
	// 向切片添加一个元素
	numbers = append(numbers, 0)
	printSliceInt("numbers", numbers)
	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSliceInt("numbers", numbers)
	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	// numbers2 = append(numbers2, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	printSliceInt("numbers", numbers)
	// 这里cap=6, 因为添加3个，容量不够，自动扩容。
	// 下面注释是知乎上搜到的一个解释，但看go的源码貌似是 < 1024 时容量每次翻倍，>= 1024 每次加 1/4
	// 经过实验，当前版本cap会增加最小满足放下添加多个元素的2的n次方个。

	// 创建切片是numbers切片的两倍容量
	numbersNew := make([]int, len(numbers), (cap(numbers))*2)
	// 拷贝 numbers => numbersNew，只会复制能放下的容量。
	copy(numbersNew, numbers)
	printSliceInt("numbersNew", numbersNew)

	// 作者：知乎用户
	// 链接：https://www.zhihu.com/question/27161493/answer/35490973

	// append的实现原理：本质上append是在原有空间中添加，若空间不足时，
	// 采用 newSlice := make([]int, len(slice), 2*(len(slice)+1))的方式进行扩容。
	// 在空间不足的情况下，append在空间扩展之后，通过copy，将原有的slice拷贝到了新的newSlice中。
	// 因此，对扩容时，会有一个内存地址变化。但是如果在满足空间大小时，内存地址不会发生变化。

	// 一个容易注意不到的情况，代码如下：
	s := []int{5}
	s = append(s, 7)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
	s = append(s, 9)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
	x := append(s, 11)
	fmt.Println("cap(x) =", cap(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
	y := append(s, 12)
	fmt.Println("cap(y) =", cap(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])
	// 注意：y操作覆盖了x操作的最后一个元素，请看打印出来的值。
	fmt.Println(s, x, y)

	// 链接：https://www.zhihu.com/question/27161493/answer/35485751
	// 1. 创建s时，cap(s) == 1，内存中数据[5]
	// 2. append(s, 7) 时，按Slice扩容机制，cap(s)翻倍 == 2，内存中数据[5,7]
	// 3. append(s, 9) 时，按Slice扩容机制，cap(s)再翻倍 == 4，内存中数据[5,7,9]，但是实际内存块容量4
	// 4. x := append(s, 11) 时，容量足够不需要扩容，内存中数据[5,7,9,11]
	// 5. y := append(s, 12) 时，容量足够不需要扩容，内存中数据[5,7,9,12]
	// 这就是后一次操作覆盖了前一次操作数据的原因。

	// 链接：https://www.zhihu.com/question/27161493/answer/147436965
	// |====================================================
	// |var             |ele        |len() |cap() |mem
	// |s               |[5]        |1     |1     |0x1234
	// |s(append(s,7))  |[5,7]      |2     |2     |0xaaaa(*)
	// |s(append(s,9))  |[5,7,9]    |3     |4     |0x8888(*)
	// |x(append(s,11)  |[5,7,9,11] |4     |4     |0x8888
	// |y(append(s,12)) |[5,7,9,12] |4     |4     |0x8888
	// |====================================================
	// 最后打印输出的时候，由于x，y引用的是同一个数组，所以结果一样
}

// 打印slice长度，容量，内存地址
func printSliceInt(name string, x []int) {
	if x != nil {
		fmt.Println(name, "len =", len(x), "cap =", cap(x), "addr =", &x[0])
	}
}

// 打印slice长度，容量，内存地址
func printSliceByte(name string, x []byte) {
	if x != nil {
		fmt.Println(name, "len =", len(x), "cap =", cap(x), "addr =", &x[0])
	}
}
