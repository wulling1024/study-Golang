package channel

import (
	"fmt"
	"testing"
)

func Test_Demo(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	fmt.Println("测试实现")
	//in := gen(done, 2, 3)
}

func gen(number ...int) <-chan int {
	out := make(chan int, len(number))
	defer close(out)
	for _, n := range number {
		out <- n
	}
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
