package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 貌似没有让出CPU时间片
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // 设为1 runtime.Gosched()才会有用
	go say("world")       //开一个新的Goroutines执行
	say("hello")          //当前Goroutines执行
	time.Sleep(1 * time.Second)
}
