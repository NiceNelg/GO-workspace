package main

import (
	"fmt"
)

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channl数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	//数字通讯
	ch := make(chan int)
	//程序是否结束
	quit := make(chan bool)

	//消费者
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		//可停止
		quit <- true
	}()

	//生成者
	fibonacci(ch, quit)
}
