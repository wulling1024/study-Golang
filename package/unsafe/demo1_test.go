package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_demo1(t *testing.T) {
	fmt.Println(unsafe.Sizeof(""))

	fmt.Println(unsafe.Alignof(12))

	//fmt.Println(unsafe.Offsetof(10))
}
