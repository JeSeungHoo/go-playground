package main

import "fmt"

// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument

func main() {
	fmt.Println(bar(foo, 1))
}

func foo(x int) int {
	x++
	return x
}

func bar(f func(x int) int, x int) int {
	y := f(x)
	y *= y
	return y
}
