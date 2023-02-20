package main

import "fmt"

func main() {
	fmt.Println("体脂计算器")
	for {
		var name string
		fmt.Println("请输入姓名：")
		fmt.Scanln(&name)

		var kg float64
		fmt.Println("请输入体重：")
		fmt.Scanln(&kg)

		var m float64
		fmt.Println("请输身高：")
		fmt.Scanln(&m)

		var age int
		fmt.Println("请输年龄：")
		fmt.Scanln(&age)

		var sex int
		fmt.Println("请输性别(1为男，0为女)：")
		fmt.Scanln(&sex)
		BMI := 1.2*(kg/(m*m)) + (0.23 * float64(age)) - 5.4 - (10.8 * float64(sex))
		fmt.Println(BMI)
		BM := float64(82) / (1.72 * 1.72)
		fmt.Println(BM)
		var p string
		fmt.Println("是否输入下一个：请输入Y/N")
		fmt.Scanln(&p)
		if p != "Y" {
			break
		}
	}

}
