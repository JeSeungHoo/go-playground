// ● write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func(i int) {
			for j := 1; j <= 10; j++ {
				c <- i * j
			}
			if i == 10 {
				close(c)
			}
		}(i)
	}

	for v := range c {
		fmt.Println(v)
	}
}
