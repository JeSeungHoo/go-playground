// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”. If you need a hint, here is one.

package main

import "fmt"

type customErr struct {
	message string
}

func (c customErr) Error() string {
	return c.message
}

func foo(e error) {
	fmt.Println(e.Error())
}

func main() {
	c := customErr{
		message: "message!!",
	}
	foo(c)
}
