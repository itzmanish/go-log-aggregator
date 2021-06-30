package queue

import (
	"sync"
	"time"

	"github.com/itzmanish/go-log-aggregator/internal/client"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
)

type Queue interface {
	Pop(key interface{})
	Push(value interface{})
	Get(key interface{}) (interface{}, bool)
	Length() int
	String() string
}

type memQueue struct {
	queue    sync.Map
	length   int
	interval time.Duration
	client   client.Client
}

func (mq *memQueue) handle() int {
	for {
		<-time.After(mq.interval)
		logger.Debug("Queue [status]Total: ", mq.length)
		go func() {
			mq.queue.Range(func(key, value interface{}) bool {
				logger.Info(key, value)
				err := mq.client.Send(value)
				if err == nil {
					mq.Pop(key)
				}
				return true
			})
		}()
	}
}

func (mq *memQueue) Length() int {
	return mq.length
}

func (mq *memQueue) Push(data interface{}) {
	mq.length++
	mq.queue.Store(mq.length, data)
}

func (mq *memQueue) Pop(key interface{}) {
	mq.queue.Delete(key)
	if mq.length == 0 {
		return
	}
	mq.length--
}

func (mq *memQueue) Get(k interface{}) (interface{}, bool) {
	return mq.queue.Load(k)
}

func (mq *memQueue) String() string {
	return "Memory queue"
}

func NewQueue(c client.Client, interval time.Duration) Queue {
	q := &memQueue{
		client:   c,
		interval: interval,
	}
	go q.handle()
	return q
}
