package main

import "fmt"

func main() {
	//arr := [5]int{1, 2, 3, 4, 5}
	//var sp = arr[1:4]
	//sp = append(sp, 6)
	//fmt.Println(arr)
	//fmt.Println(sp)
	//fmt.Println(reflect.TypeOf(arr))
	//fmt.Println(reflect.TypeOf(sp))
	//var xp = [3]string{"小明", "23", "4年级"}
	//var xb = [3]string{"小碧", "25", "2年级"}
	//var xd = [3]string{"小迪", "26", "3年级"}
	//
	//person := [...][3]string{
	//	xp, xb, xd,
	//}
	//for index, value := range person {
	//	print(index)
	//	for _, val := range value {
	//		print(index, val)
	//	}
	//	print("\n")
	//}

	var year int
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	var m = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var d [12][31]int
	for idx, val := range m {
		switch val {
		case 1, 3, 5, 7, 8, 10, 12:
			for i := 0; i < 31; i++ {
				d[idx][i] = i + 1
			}
			break
		case 2:
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				for i := 0; i < 28; i++ {
					d[idx][i] = i + 1
				}
			} else {
				for i := 0; i < 29; i++ {
					d[idx][i] = i + 1
				}
			}
			break
		case 4, 6, 9, 11:
			for i := 0; i < 30; i++ {
				d[idx][i] = i + 1
			}
			break
		}
	}
	fmt.Println(d)
}

func month(m [12]int, year int) [12][31]int {
	var d [12][31]int
	for idx, val := range m {
		switch val {
		case 1, 3, 5, 7, 8, 10, 12:
			for i := 0; i < 31; i++ {
				d[idx][i] = i + 1
			}
			break
		case 2:
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				for i := 0; i < 28; i++ {
					d[idx][i] = i + 1
				}
			} else {
				for i := 0; i < 29; i++ {
					d[idx][i] = i + 1
				}
			}
			break
		case 4, 6, 9, 11:
			for i := 0; i < 30; i++ {
				d[idx][i] = i + 1
			}
			break
		}
	}
	return d
}
