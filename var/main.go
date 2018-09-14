package main

import "fmt"

/**
 *
 */
func main() {
	//常规命名规则：var 变量名 变量类型
	var name string
	name = "tom"
	fmt.Println("name is", name)

	//忽略变量类型的变量命名规则：var 变量=值 会根据值的类型隐式给变量声明类型
	var name1 = "Tom"
	fmt.Println("name is", name1)
}
