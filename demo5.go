package main

import (
	"fmt"
)

func addNum(a, b int) int {
	return a + b
}

//func init() {
//	fmt.Println("这是初始化")
//}

func foo() (int, string) {
	return 10, "这是一个字符串"
}

func main() {
	//c := addNum(1, 3)
	//fmt.Println(c)
	//buf := make([]byte, 1024)
	//f, _ := os.Open("./index.html")
	//defer f.Close()
	//for {
	//	n, _ := f.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	os.Stdout.Write(buf[:n])
	//}

	a, b := foo()
	fmt.Println(a)
	fmt.Println(b)

	s2 := "博客"
	runes2 := []rune(s2)
	runes2[0] = '狗'
	fmt.Println(string(runes2))

	var arr = [5]int{1, 2, 3, 4}
	for index, value := range arr {
		fmt.Printf("%d: %d ", index, value)
	}

	type Person struct {
		name string
		age  int
	}

	var person = [...]Person{{name: s2, age: a}}
	fmt.Println(person)
}
