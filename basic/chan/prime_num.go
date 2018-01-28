package main

import "fmt"

// 将序列 2, 3, 4, … 发送至信道'ch'。
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // 将 'i' 发送至信道'ch'。
	}
}

// 将值从信道'src'中复制至信道'dst'，
// 移除可被'prime'整除的数。
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // 循环遍历从'src'接收的值。
		if i%prime != 0 {
			dst <- i // 将'i'发送至'dst'。
		}
	}
}

// 质数筛：将过滤器串联在一起处理。
func sieve() {
	ch := make(chan int) // 创建一个新信道。
	go generate(ch)      // 将generate()作为子进程开始。
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
