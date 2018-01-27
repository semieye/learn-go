package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	fmt.Println(rand.Float64())
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// 通过在boring的循环中增加一个select，在main中我们便可以通过向quit 写入数据的方式来控制boring的退出。
func boringWithQuit(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				return
			}
		}
	}()
	return c
}

// goroutines 通过 chan 进行通信同步
func main() {
	defer fmt.Println("You're both boring; I'm leaving.")
	joe := boring("Joe say:")
	ann := boring("Ann say:")
	for i := 0; i < 1; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	c := multipleGoroutines(boring("MMM"), boring("UUU"))
	for i := 0; i < 1; i++ {
		fmt.Println(<-c)
	}
	c = selectGoroutines(boring("SSS"), boring("LLL"))
	for i := 0; i < 1; i++ {
		fmt.Println(<-c)
	}
	quit := make(chan bool)
	c = boringWithQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

// 多路复用
func multipleGoroutines(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// select方式
func selectGoroutines(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
