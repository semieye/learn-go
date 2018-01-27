package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		ci := make(chan int)           // 整数类型的无缓冲信道
		cj := make(chan int, 0)        // 整数类型的无缓冲信道
		cs := make(chan *os.File, 100) // 指向文件指针的带缓冲信道
	*/

	c := make(chan int) // 分配一个信道
	// 在Go程中启动排序。当它完成后，在信道上发送信号。
	go func() {
		// list.Sort()
		sleepAWhile()
		c <- 1 // 发送信号，什么值无所谓。
	}()
	// doSomethingForAWhile()
	<-c // 等待排序结束，丢弃发来的值。
}

func sleepAWhile() {
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	//例如，rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	fmt.Print(rand.Intn(100), ",")
	//rand.Float64 返回一个64位浮点数 f，0.0 <= f <= 1.0。
	fmt.Println(rand.Float64())

	t := time.Duration(rand.Intn(1e3)) * time.Millisecond
	fmt.Println("sleep awhile:", t)
	time.Sleep(t)
}
