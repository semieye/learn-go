package main

import (
	"fmt"
	"os"
)

func main() {
	// 目录操作
	os.Mkdir("test", 0777)
	os.MkdirAll("test/test1/test2", 0777)
	err := os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("test")

	// 写文件
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	// 读文件
	userFile2 := "test.txt"
	fl, err := os.Open(userFile2)
	if err != nil {
		fmt.Println(userFile2, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	// 删除文件
	// func Remove(name string) Error
}
