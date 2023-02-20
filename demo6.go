package main

import (
	"fmt"
	"time"
)

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]
var slice2 []int = arr[5:10]
var slice3 []int = arr[0:len(arr)]
var slice4 = arr[:len(arr)-1]

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}
func main() {
	//map1 := make(map[int]string, 5)
	//map1[1] = "www.topgoer.com"
	//map1[2] = "rpc.topgoer.com"
	//map1[5] = "ceshi"
	//map1[3] = "xiaohong"
	//map1[4] = "xiaohuang"
	//sli := []int{}
	//fmt.Println(map1)
	//for k, _ := range map1 {
	//	println(k)
	//	sli = append(sli, k)
	//}
	//fmt.Println(sli)
	//sort.Ints(sli)
	//fmt.Println(sli)
	//for i := 0; i < len(map1); i++ {
	//	fmt.Println(map1[sli[i]])
	//}
	//var a int = 100
	//if a < 20 {
	//	fmt.Println("a小于20")
	//} else {
	//	fmt.Println("a大于20")
	//}
	//var c1, c2, c3 chan int
	//var i1, i2 int
	//select {
	//case i1 = <-c1:
	//	fmt.Printf("received ", i1, " from c1\n")
	//case c2 <- i2:
	//	fmt.Printf("sent ", i2, " to c2\n")
	//case i3, ok := (<-c3): // same as: i3, ok := <-c3
	//	if ok {
	//		fmt.Printf("received ", i3, " from c3\n")
	//	} else {
	//		fmt.Printf("c3 is closed\n")
	//	}
	//default:
	//	fmt.Printf("no communication\n")
	//}

	test()

}

var resChan = make(chan int)

// do request
func test() {
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

func doData(data int) int {
	//...
	fmt.Println(data)
	return 1
}
