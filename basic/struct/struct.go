package main

import "fmt"

// Books the struct of books
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

// Human a struct
type Human struct {
	name   string
	age    int
	weight int
	phone  string //与emp中同名字段
}

// Student a struct include Human
type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

// Skills a slice
type Skills []string

// Employee struct
type Employee struct {
	Human      // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookID = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookID = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	printStudent()

	printEmp()
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookID : %d\n", book.bookID)
}

func printStudent() {
	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100, "xxxx"}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

func printEmp() {
	Bob := Employee{Human{"Bob", 34, 10, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
	//最外层的优先访问，也就是当你通过Employee.phone访问的时候，是访问Employee里面的字段，而不是human里面的字段
}
