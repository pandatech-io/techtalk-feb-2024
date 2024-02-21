package graceful

import (
	"fmt"
	"sync"
	"time"
)

func RunGoroutine() {
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func RunChannel() {
	ch := make(chan string)

	go func() {
		ch <- "hai from goroutine"
	}()

	fmt.Println(<-ch)
}

func RunBufferChannel() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			ch <- num
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
	for num := range ch {
		fmt.Println(num)
	}
}
