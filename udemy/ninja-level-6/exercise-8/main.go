package main

import "fmt"

// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func
func main() {
	x := maker()
	y := x()
	fmt.Println(y)
}

func maker() func() int {
	return func() int {
		return 5
	}
}
