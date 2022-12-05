package main

import "fmt"

func main() {
	// 정수를 넣는 채널
	c := make(chan int)

	// 채널 c에 42라는 값을 넣음
	go func() {
		c <- 42
	}()

	// 채널 c에서 값을 가져옴
	fmt.Println(<-c)
}
