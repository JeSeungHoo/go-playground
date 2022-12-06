package main

import (
	"fmt"

	"gitgub.com/JeSeungHoo/go-playground/udemy/ninja-level-13/exercise-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
