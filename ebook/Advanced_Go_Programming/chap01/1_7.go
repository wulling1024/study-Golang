package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan int, 64)

	go Producer(2, ch)
	go Producer(3, ch)

	go Consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("Quit (%v)\n", <-sig)
}

// 生产者
func Producer(factory int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factory
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
