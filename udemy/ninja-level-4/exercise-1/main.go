package main

import "fmt"

func main() {
	// 합성 리터럴로 배열 만들기 -- go에서는 슬라이스를 써라!
	array := [5]int{10, 11, 12, 13, 14}
	// for range
	for i, v := range array {
		fmt.Println(i, v)
	}
	// type
	fmt.Printf("%T", array)
}
