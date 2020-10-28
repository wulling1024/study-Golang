package main

import "fmt"

// 参数类型(chan int): 即可读，又可写
func ReadAndWrite(ch chan int) chan int {
	a := <-ch
	fmt.Printf("(chan int) 读写管道内容: %v \n", a)
	ch <- a * 10

	return ch
}

// 参数类型(<-chan int): 只可读
func Read(ch <-chan int) <-chan int {
	a := <-ch
	fmt.Printf("(chan int) 读写管道内容: %v \n", a)
	//ch <- a*10 // Invalid operation: ch <- a*10 (send to receive-only type <-chan int)

	return ch
}

// 参数类型(chan<- int): 只可写
func Write(ch chan<- int) chan<- int {
	var a int
	//a = <- ch //Invalid operation: <- ch (receive from send-only type chan<- int)
	fmt.Printf("(chan int) 读写管道内容: %v \n", a)
	ch <- a * 10

	return ch
}

func main() {

}
