package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/**
该测试的目的：
	1、判断goroutine中基本类型、slice、map是值传递还是引用传递
*/
type Demo struct {
	x int
	y []int
	z map[int]string
}

func Test20210527_1(t *testing.T) {
	var wg sync.WaitGroup

	demo := new(Demo)

	var a int = 1
	demo.x = 1
	var b []int
	//demo.y = []int{1, 2}
	c := make(map[int]string)
	c[1] = "one"
	demo.z = c
	wg.Add(1)
	go func(x int, y []int, z map[int]string) {
		defer wg.Done()
		x = x * 2
		//for index, value := range y {
		//	y[index] = value * 2
		//}
		y = append(y, x)
		z[1] = "update one"
	}(a, b, c)
	wg.Add(1)
	go func(t *Demo) {
		defer wg.Done()
		t.x = t.x * 2
		//for _, value := range t.y {
		//	t.y = append(t.y, value + 2)
		//}
		t.y = append(t.y, 2)

		t.z[1] = "update one"
	}(demo)

	wg.Wait()
	fmt.Println(a)
	fmt.Println("=============================")
	fmt.Println(b)
	fmt.Println("=============================")
	fmt.Println(c[1])
	fmt.Println("=============================")
	fmt.Println(*demo)
}
