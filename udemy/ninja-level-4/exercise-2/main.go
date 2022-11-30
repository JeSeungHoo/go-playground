package main

import "fmt"

func main() {
	// 슬라이스
	slice := []int{10, 11, 12, 13, 14}
	// for range
	for i, v := range slice {
		fmt.Println(i, v)
	}
	// type
	fmt.Printf("%T", slice)
}
