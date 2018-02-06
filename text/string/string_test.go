package demo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var _ = &testing.T{}

// 示例
func Example1() {
	// func Contains(s, substr string) bool
	// 字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	//Output:
	//true
	//false
	//true
	//true
}
func Example2() {
	// func Join(a []string, sep string) string
	// 字符串链接，把slice a通过sep链接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// func Index(s, sep string) int
	// 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	//Output:foo, bar, baz
	//4
	//-1
}
func Example3() {
	// func Repeat(s string, count int) string
	// 重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	//Output:banana
}

func Example4() {
	// func Replace(s, old, new string, n int) string
	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//Output:oinky oinky oink
	//moo moo moo
}

func Example5() {
	// func Split(s, sep string) []string
	// 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//Output:["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]
}

func Example6() {
	// func Trim(s string, cutset string) string
	// 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	//Output:["Achtung"]
}
func Example7() {
	// func Fields(s string) []string
	// 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//Output:Fields are: ["foo" "bar" "baz"]
}

func Example8() {
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

func Example9() {
	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}

func Example10() {
	// Parse 系列函数把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
