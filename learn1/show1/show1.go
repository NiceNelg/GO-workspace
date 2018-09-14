package show1

import "fmt"
import "learn1/show2"

func init() {
	fmt.Println("show1")
}

func Show1() {
	show2.Show2()
}
