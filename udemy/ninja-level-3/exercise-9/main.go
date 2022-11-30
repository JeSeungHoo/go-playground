package main

import "fmt"

func main() {
	favSport := "농구"
	switch favSport {
	case "축구":
		fmt.Println("축구가 최고지~")
	case "야구":
		fmt.Println("야구가 재밌어!")
	case "농구":
		fmt.Println(" M J~")
	case "씨름":
		fmt.Println("천하장사")
	}
}
