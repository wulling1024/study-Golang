package goroutine

import (
	"fmt"
	"testing"
)

// 开启100个goroutine,查看是否会出现并发问题
func TestSliceAppend(t *testing.T) {

	rsp := make([]int, 0)

	for i := 0; i < 100; i++ {
		go func(num int) {
			rsp = append(rsp, i)
		}(i)
	}

	fmt.Println(len(rsp))
}
