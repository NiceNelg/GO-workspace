package main

import "fmt"

func main() {
	//常规数组声明规则：变量名 [数组长度]变量类型
	var array [10]int
	fmt.Println(array)

	//数组声明并赋值
	var array1 = [6]string{"1", "2", "3", "4", "5", "6"}
	fmt.Println("", array1[0])

	//不声明数组的长度，数组的长度由元素个数决定，注意中括号中的省略号必须要写，否则会变成切片类型
	var array2 = [...]int{1, 2, 3, 4}
	fmt.Println(array2)
	fmt.Println("", len(array2))
}
