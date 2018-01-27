package main

import (
	"fmt"
	"math/rand"
	"time"
)

var syn chan int = make(chan int)

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("i am running!")
	}
	syn <- 1
}

func init() {
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	fmt.Println(rand.Float64())
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// func main() {
// 	go foo()
// 	<-syn
// }

// goroutines 通过 chan 进行通信同步
func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
