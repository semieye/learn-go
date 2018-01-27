// 切片：动态数组
// 包注释：会被godoc读取生成文档

package main

import (
	"fmt"
)

func main() {
	// type Transform [3][3]float64 // 一个 3x3 的数组，其实是包含多个数组的一个数组。
	// type LinesOfText [][]byte    // 包含多个字节切片的一个切片。

	// text := LinesOfText{
	// 	[]byte("Now is the time"),
	// 	[]byte("for all good gophers"),
	// 	[]byte("to bring some fun to the party."),
	// }

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)

}

// Append func from the effective_go document
func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // 重新分配
		// 为了后面的增长，需分配两份。
		newSlice := make([]byte, (l+len(data))*2)
		// copy 函数是预声明的，且可用于任何切片类型。
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

// 二维切片 两种分配方法
/*
// 分配顶层切片。
picture := make([][]uint8, YSize) // 每 y 个单元一行。
// 遍历行，为每一行都分配切片
for i := range picture {
	picture[i] = make([]uint8, XSize)
}
*/
/*
现在是一次分配，对行进行切片：
// 分配顶层切片，和前面一样。
picture := make([][]uint8, YSize) // 每 y 个单元一行。
// 分配一个大的切片来保存所有像素
pixels := make([]uint8, XSize*YSize) // 拥有类型 []uint8，尽管图片是 [][]uint8.
// 遍历行，从剩余像素切片的前面切出每行来。
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
*/

// 内建函数 append 说明
// func append(slice []T, 元素 ...T) []T
