package pubsub

import (
	"sync"
	"time"
)

type (
	// 订阅者为一个通道
	subscriber chan interface{}
	// 主题为一个过滤器
	topicFunc func(v interface{}) bool
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex             //读写锁
	buffer      int                      //订阅队列的缓存大小
	timeout     time.Duration            //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

// 构建一个发布者对象：队列大小、超时时间
func NewPublisher(buffer int, time time.Duration) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     time,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个订阅者，订阅所有的主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscriberTopic(nil)
}

// 添加一个订阅者，订阅过滤器筛选过的主题
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 发送主题，容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// 关闭发布者对象，同时关闭所有的订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
