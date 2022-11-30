package main

import "fmt"

func main() {
	birth := 1996
	for {
		if birth > 2022 {
			break
		}
		fmt.Println(birth)
		birth++
	}
}
