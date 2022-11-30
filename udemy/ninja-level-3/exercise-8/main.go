package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false")
	case 4 == 5:
		fmt.Println("4==5")
	case 1 == 1:
		fmt.Println("1==1")
	}
}
