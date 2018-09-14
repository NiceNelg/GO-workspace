package main

import (
	"defer/test"
	"fmt"
)

func main() {
	var i int
	test.Test()
	for i = 0; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
