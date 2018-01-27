package main

import (
	"fmt"
)

type phone interface {
	call()
}

type nokiaPhone struct {
}

func (n nokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type iPhone struct {
}

func (i iPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var p phone

	p = new(nokiaPhone)
	p.call()

	p = new(iPhone)
	p.call()

}
