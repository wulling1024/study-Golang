package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test20210530(t *testing.T) {
	upload()
	uploadV1()
}

func upload() {

	data := make([]int, 0)
	start := time.Now()
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	ch := make(chan error, 4)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(tag int) {
			defer wg.Done()
			fileManager()
			fmt.Printf("协程编号：%v \n", tag)
			//if tag == 2 || tag == 3 {
			//	Err := fmt.Errorf("error massage: %v", tag)
			//	ch <- Err
			//	return
			//}
			mutex.Lock()
			//d = append(d, tag)
			data = append(data, tag)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	close(ch)
	if len(ch) != 0 {
		fmt.Println(<-ch)
		return
	}

	fmt.Printf("并发执行的耗时：%v \n", time.Now().Sub(start))
	fmt.Printf("数据信息：%v \n", data)
}
func uploadV1() {
	start := time.Now()
	for i := 0; i < 4; i++ {
		fileManager()
	}

	fmt.Printf("串行执行的耗时：%v \n", time.Now().Sub(start))
}

func fileManager() {
	// waste 2s
	time.Sleep(2 * time.Second)
}
