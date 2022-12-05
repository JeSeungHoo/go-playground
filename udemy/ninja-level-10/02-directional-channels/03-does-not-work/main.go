package main

import "fmt"

func main() {
	// 수신 전용 채널
	c := make(<-chan int)

	go func() {
		//송신 불가
		// c <- 42
	}()

	fmt.Println(<-c)
}
