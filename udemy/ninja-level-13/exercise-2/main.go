package main

import (
	"fmt"

	"gitgub.com/JeSeungHoo/go-playground/udemy/ninja-level-13/exercise-2/quote"
	"gitgub.com/JeSeungHoo/go-playground/udemy/ninja-level-13/exercise-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
