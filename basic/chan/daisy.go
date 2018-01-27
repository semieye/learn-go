// 菊花链(Daisy-chain)

package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// 代码初始化了100000个channel，并把他们按照顺序连接起来。
// 最后向最右边的channel写入一个数据，从最左边的channel读出来。
// 这种菊花链的模型非常适合作为过滤器filter来使用，通过channel来连接filter会显得十分方便。
func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
