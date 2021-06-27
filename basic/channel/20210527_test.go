package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test20210527(t *testing.T) {

	var wg sync.WaitGroup

	ch := make(chan error, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i == 2 {
				return
			}
			e := fmt.Errorf("demo error %v", i)
			ch <- e
		}(i)
	}
	wg.Wait()
	close(ch)
	if len(ch) != 0 {
		fmt.Println(len(ch))
		<-ch
		fmt.Println(<-ch)
	}
	ch1 := make(chan int)
	fmt.Println(len(ch1))
}
