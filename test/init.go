package test

import "fmt"

// 测试 import _ 导入包仅仅为了执行init函数
func init() {
	fmt.Println("init() come here.")
}
