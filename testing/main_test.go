package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	var j int
	for i := 0; i < 1000; i++ {
		j++
	}
	fmt.Println(j)
	if j != 1000 {
		t.Error("Wrong Print")
	}
}

func testPrint2(t *testing.T) {
	var j int
	for i := 0; i < 10000; i++ {
		j++
	}
	fmt.Println(j)
	if j != 10000 {
		t.Error("Wrong Print2")
	}
}

func TestAll(t *testing.T) {
	t.Run("First", testPrint)
	t.Run("Second", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Test begin...")
	m.Run()
	fmt.Println("Test end.")
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		aaa(n)
	}
}

func aaa(n int) int {
	return n
}
