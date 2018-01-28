package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	// 无缓冲的
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// 有缓冲的
	d := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	// if buff = 1, fatal error: all goroutines are asleep - deadlock!
	d <- 1
	d <- 2
	fmt.Println(<-d)
	fmt.Println(<-d)
}

// 默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
// 这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
// 所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
// 其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
// 无缓冲channel是在多个goroutine之间同步很棒的工具。
