package goroutine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
该测试的目的：
	1、一个函数内部开启了多个goroutine, 存在一个goroutine出错，整个函数退出了。
      那其它还未执行结束的goroutine是否还会继续执行？

测试结果说明：
	1、只要主程序没有截止，goroutine就会运行直到该goroutine运行结束
*/

func Test20210527(t *testing.T) {
	qiuqiu()
	time.Sleep(time.Duration(6 * 1000 * 1000 * 1000))
}

func qiuqiu() {
	ch := make(chan error)
	for i := 0; i < 10; i++ {
		go func(key int) {
			err := gg(key)
			if err != nil {
				ch <- err
				return
			}
		}(i)
	}
	Err := <-ch
	if Err != nil {
		fmt.Println("func qiuqiu exit ...")
		return
	}
}

func gg(i int) error {
	if i != 1 {
		time.Sleep(time.Duration(rand.Intn(10) * 1000 * 1000 * 1000))
		fmt.Println(i)
		return nil
	}
	return fmt.Errorf("at %d err", i)
}
