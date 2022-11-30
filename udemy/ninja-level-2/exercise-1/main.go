package main

import "fmt"

func main() {
	a := 16
	// 10진수 2진수 2진수 16진수 16진수
	// 16 10000 0b10000 10 0x10
	fmt.Printf("%d %b %#b %x %#x", a, a, a, a, a)
}
