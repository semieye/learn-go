package demo

import "testing"
import "fmt"

var _ = &testing.T{}

func Foo() {
	fmt.Println("Hello")
}

func ExampleFoo() {
	Foo()
	// Output:
	// Hello
	//
}

// 普通单元测试用例
func TestFoo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// ...
}

// 性能测试用例
func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// ...
	}
}

// 示例
func ExampleBaz() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleQux() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
