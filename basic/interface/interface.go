package main

import (
	"fmt"
	"sort"
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

	var t interface{}
	t = 0.4325345
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 输出 t 是什么类型
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	}

	//用Sequence实现sort.Interface的方法，实现排序后打印
	seq := Sequence{2, 1, 0, 3, 9, 7, 6, 5, 8, 4}
	fmt.Println(seq.String())

}

// Sequence interface
type Sequence []int

// Methods required by sort.Interface.
// sort.Interface 所需的方法。
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
// 用于打印的方法 - 在打印前对元素进行排序。
func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
