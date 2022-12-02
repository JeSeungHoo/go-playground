// in addition to the main goroutine, launch two additional goroutines
// ○ each additional goroutine should print something out
// ● use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPUs", runtime.NumCPU())
	fmt.Println("begin Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("something1")
		wg.Done()
	}()
	go func() {
		fmt.Println("something2")
		wg.Done()
	}()

	fmt.Println("mid CPUs", runtime.NumCPU())
	fmt.Println("mid Goroutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("end CPUs", runtime.NumCPU())
	fmt.Println("end Goroutine", runtime.NumGoroutine())
	fmt.Println("done")
}
