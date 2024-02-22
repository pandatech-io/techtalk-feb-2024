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
	capacity := 10
	ch := make(chan int, capacity)

	var wg sync.WaitGroup

	wg.Add(capacity) // 10
	for i := 0; i < 10; i++ {
		go func(num int) {
			ch <- num
			wg.Done() // 9 8 7 ... 0
		}(i)
	}

	wg.Wait()
	close(ch)
	for num := range ch {
		fmt.Println(num)
	}
}
