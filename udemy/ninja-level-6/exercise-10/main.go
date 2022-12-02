package main

import "fmt"

// Closure is when we have “enclosed” the scope of a variable in some code block. For this
// hands-on exercise, create a func which “encloses” the scope of a variable:

func main() {
	counter1 := foo()
	counter2 := foo()
	counter1()
	counter1()
	counter1()
	counter1()
	counter1()

	counter2()
	counter2()
	counter2()

	fmt.Println(counter1())
	fmt.Println(counter2())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
