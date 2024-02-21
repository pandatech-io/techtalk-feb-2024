package graceful

import (
	"fmt"
	"sync"
	"time"
)

func RunNotGrace() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the first loop")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the second loop")
		}
	}()
	// Wait for ongoing process to finish
	wg.Wait()
	fmt.Println("Process cleanup...") // This won’t be called
}
