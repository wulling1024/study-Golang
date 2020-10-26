package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 100

	fmt.Printf("a: %p \n", &a)
	fmt.Printf("b: %p \n", &b)

	// 函数参数形式
	updateA(a)
	// 闭包形式
	func(a int) {
		fmt.Printf("update b: %p \n", &b)
	}(b)

}


// 函数参数形式
func updateA(a int) {
	fmt.Printf("update a: %p \n", &a)
}
