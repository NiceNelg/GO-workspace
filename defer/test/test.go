package test

import "fmt"

func Test() int {
	var i int
	for i = 10; i < 20; i++ {
		defer fmt.Printf("%d\n", i)
	}
	return i
}
