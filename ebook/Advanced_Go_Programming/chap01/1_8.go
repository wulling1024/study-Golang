package main

import (
	"fmt"
	"github.com/wulling1024/study-Golang/ebook/Advanced_Go_Programming/chap01/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(10, 100*time.Millisecond)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscriberTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world")
	p.Publish("hello, golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 主协程运行一段时间后退出
	time.Sleep(3 * time.Second)
}
