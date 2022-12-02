package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func main() {
	foo()
}

func foo() {
	defer bar()
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
