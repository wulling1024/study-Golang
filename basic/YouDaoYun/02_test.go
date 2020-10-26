package main

import (
	"fmt"
	"testing"
)

// 测试运算符
func Test_Main_02(t *testing.T) {
	fmt.Println(12&4)  // 4
	fmt.Println(12^4)  // 8
	fmt.Println(12|4)  // 12
}
