package main

import "fmt"

func main() {
	fmt.Println(a())
}

func a() (v int) {
	// 延迟函数 defer
	defer func() {
		v++
	}()
	return 43
}
