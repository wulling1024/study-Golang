package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only Once")
	}

	ch := make(chan bool, 0)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			ch <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}
