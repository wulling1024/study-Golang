package function

import (
	"fmt"
	"testing"
)

/**
验证问题：
	1、函数参数的 []int 类型到底是 切片类型 还是 数组类型

结论：
	1、数组的类型 [2]int 本身就是一个具体的类型，例如 [4]int
	2、函数参数的 []int 显示的就是切片类型，如果是数组类型，必须为具体的数组类型，比如 func solve(nums [2]int)
 */

func Test_demo(t *testing.T) {
	var a [2]int = [2]int{1,2}
	solve2(a)

	var b []int = []int{1,2}
	solve(b)
}

func solve(num []int) {
	var a = num
	a[0] = a[0] * 2
	fmt.Println(num)
	fmt.Println(a)
}

func solve2(num [2]int) {
	var a = num
	a[0] = a[0] * 2
	fmt.Println(num)
	fmt.Println(a)
}
