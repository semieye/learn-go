package init

import (
	"fmt"
)

// 测试 import _ 导入包仅仅为了执行init函数

// 当前实现提供了几个在引导过程中有用的内建函数。这些函数因完整性而被保留， 但不保证会继续留在该语言中。它们并不返回结果。
// print      打印所有实参；实参的格式取决于具体实现
// println    类似print，但会在实参之间打印空格并在末尾打印新行

func init() {
	println("println")
	print("os.Args")
	fmt.Println("init() come here.")
}
