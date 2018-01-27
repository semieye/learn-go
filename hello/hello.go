//my first go program.
package main

import (
	"fmt"
	"log"
	"os"
	// . "fmt" // . 操作可以使用包函数不需要前缀 例如直接 Printf
	// f "fmt" // 别名 f.Printf

	_ "github.com/semieye/learn-go/test" //导入该包仅为了执行其中init函数
	"github.com/semieye/learn-go/utils"
)

//	要让编译器停止关于未使用导入的抱怨，需要空白标识符来引用已导入包中的符号。下面 fd 变量也是类似。
var _ = os.Args

func main() {
	fmt.Printf("Hello, world. \n")
	fmt.Println(utils.Reverse("!selpmaxe oG ,olleH"))
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")

	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd.
	_ = fd
}
