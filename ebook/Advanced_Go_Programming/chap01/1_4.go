package main

import "fmt"

/*
a 与 s1 指向同一个底层数组
make 函数自建底层数组
*/

func main() {
	a := []int{1,2,3,4}
	s1 := a[:0]
	s2 := make([]int, 0)

	fmt.Printf("a: %p \n", &a)
	fmt.Printf("s1: %p \n", &s1)
	fmt.Printf("s2: %p \n", &s2)

	for _, value := range a {
		s1 = append(s1, value * 2)
		s2 = append(s2, value + 2)
	}

	for _, value := range a {
		fmt.Println(value)
	}

	for _, value := range s1 {
		fmt.Println(value)
	}

	for _, value := range s2 {
		fmt.Println(value)
	}

}
