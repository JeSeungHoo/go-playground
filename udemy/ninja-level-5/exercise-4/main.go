package main

import "fmt"

func main() {
	// Create and use an anonymous struct.
	car := struct {
		wheel int
		color string
		cost  int
	}{
		wheel: 4,
		color: "Black",
		cost:  50000000,
	}
	fmt.Println(car)
}
