package main

import "fmt"

func main() {
	// 无法捕获异常
	//defer recover()

	// 正确的异常捕获
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("error is %v \n", msg)
		}
	}()

	panic("TODO")
}
