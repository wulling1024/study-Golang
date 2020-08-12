package YouDaoYun

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
对应 basic/05 的内容
 */
func Test_Main(t *testing.T){
	//testWaitGroup()
	//testChannel()
	testContext()
}

func testWaitGroup() {
	var wait sync.WaitGroup

	i := 0

	wait.Add(2)

	// 关于两者的信息输出顺序之分：FILO(先进后出)
	go func() {
		defer wait.Done() // finish task2
		fmt.Println("goroutine 2 done")
		i++
	}()
	go func() {
		defer wait.Done() // finish task1
		fmt.Println("goroutine 1 done")
		i++
	}()

	wait.Wait() // wait for tasks to be done
	fmt.Println("all goroutine done")
	fmt.Println(i)
}

func testChannel(){
	notify := make(chan bool)
	go func(){
		for {
			select {
			case <- notify:
				fmt.Println("exit")
				return
			case <- time.After(2 * time.Second):
				fmt.Println("monitoring")
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("notify exit")
	notify <- true
	time.Sleep(5 * time.Second)
}

func testContext(){
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx)
	fmt.Println("client release connection, need to notify A, B exit")
	time.Sleep(5 * time.Second)
	cancel() //mock client exit, and pass the signal, ctx.Done() gets the signal  time.Sleep(3 * time.Second)
	time.Sleep(3 * time.Second)
}

func bar(ctx context.Context){
	for {
		select{
		case <- ctx.Done():
			fmt.Println("B exit")
			return
		case <- time.After(2 * time.Second):
			fmt.Println("B do something")
		}
	}
}

func foo(ctx context.Context){
	go bar(ctx)

	for {
		select {
		case <- ctx.Done():
			fmt.Println("A exit")
			return
		case <- time.After(1 * time.Second):
			fmt.Println("A do something")
		}
	}
}