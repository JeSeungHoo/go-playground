package main

import "fmt"

// ● Assign a func to a variable, then call that func
func main() {
	x := func() {
		fmt.Println("Hi")
	}
	x()
}
