package main

func main() {
	// 송신 전용 채널
	c := make(chan<- int)

	go func() {
		c <- 42
	}()

	// 수신 불가
	// fmt.Println(<-c)
}
