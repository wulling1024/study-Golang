package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()

	select {
	default:
		fmt.Println("hello")
	case <-ch:
		return
	}
}

func main() {
	ch := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, ch)
	}

	time.Sleep(time.Second)
	close(ch)
	wg.Wait()
}
