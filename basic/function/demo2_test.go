package function

import (
	"fmt"
	"testing"
)

func Test_demo2(t *testing.T) {
	data := []int{1}
	a(data)
	fmt.Printf("主函数：%v \n", data)

	data1 := [1]int{1}
	a1(data1)
	fmt.Printf("主函数：%v \n", data1)

	data2 := [1]int{1}
	a2(&data2)
	fmt.Printf("主函数：%v \n", data2)
}

//参数类型为：[]int 切片类型，是引用类型，还是值传递
//根据结果可以判断，切片类型是 引用传递（内部对于元素的修改，会对外部元素产生影响）
func a(num []int) {
	num[0] = num[0] + 1
	fmt.Printf("函数内：%v \n", num)
}

// 参数类型为：[n]int 数组类型，是值传递
// 内部参数的修改，不会对外部变量造成影响
func a1(num [1]int) {
	num[0] = num[0] + 1
	fmt.Printf("函数内：%v \n", num)
}

// 参数类型为：*[n]int 指针数组类型，是引用传递
func a2(num *[1]int) {
	num[0] = num[0] + 1
	fmt.Printf("函数内：%v \n", num)
}
