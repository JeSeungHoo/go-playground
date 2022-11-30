package main

import "fmt"

func main() {
	// Create a slice of a clice of string
	one := []string{"James", "Bond", "Shaken, not stirred"}
	two := []string{"Miss", "Moneypenny", "helooooooo, James."}
	x := [][]string{one, two}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
