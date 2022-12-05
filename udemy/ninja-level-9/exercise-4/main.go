// Fix the race condition you created in the previous exercise by using a mutex
// ● it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count", count)
}

// go run -race main.go
