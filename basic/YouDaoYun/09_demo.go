package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start....")
	test02()
}

func test01() {
	var param int
	fmt.Scan(&param)
	ch := make(chan int, 1)
	ch <- param

Exit:
	for {
		select {
		case e, ok := <-ch:
			if ok {
				fmt.Sprintf("input: %d", e)
				break Exit
			}
		default:
			fmt.Println("null")
		}
	}
}

func test02() {
	// 构建done channel，整个管道里分享done，并在管道退出时关闭这个channel
	// 以此通知所有Goroutine该退出了。
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)

	// 发布sq的工作到两个都从in里读取数据的Goroutine
	c1 := sq(in)
	c2 := sq(in)

	// 处理来自output的第一个数值
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 或者 9

	// done会通过defer调用而关闭
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//gen函数启动一个Goroutine，将整数数列发送给channel，如果所有数都发送完成，关闭这个channel
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

// 从一个channel接收整数，并求整数的平方，发送给另一个channel.
// mission的循环中退出，因为我们知道如果done已经被关闭了，也会关闭上游的gen状态.
// mission通过defer语句，保证不管从哪个返回路径，它的out channel都会被关闭.

func mission(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				//case <- done:
				return
			}
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每个cs中的输入channel启动一个output Goroutine。outpu从c里复制数值直到c被关闭
	// 或者从done里接收到数值，之后output调用wg.Done
	output := func(cs <-chan int) {
		for n := range cs {
			select {
			case out <- n:
				//case <- done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个Goroutine，当所有output Goroutine都工作完后（wg.Done），关闭out，
	// 保证只关闭一次。这个Goroutine必须在wg.Add之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
