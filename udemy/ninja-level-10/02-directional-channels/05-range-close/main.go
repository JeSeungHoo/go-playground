package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("done")
}

func foo(c1 chan<- int) {
	for i := 0; i < 5; i++ {
		c1 <- i
	}
	// 채널 c 자체를 닫음
	close(c1)
}

func bar(c <-chan int) {
	// 채널이 닫힐 때까지(close) 반복
	for v := range c {
		fmt.Println(v)
	}
}

// 고루틴 생각
