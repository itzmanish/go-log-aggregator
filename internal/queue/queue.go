package queue

import (
	"sync"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/client"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
)

type Queue interface {
	Pop(key interface{})
	Push(value interface{})
	Get(key interface{}) (interface{}, bool)
	ForEach(f func(key interface{}, value interface{}) bool)
	Length() int
	String() string
}

type memQueue struct {
	queue  sync.Map
	length int
	client client.Client
}

func (mq *memQueue) handle() int {
	for {
		<-time.After(10 * time.Second)
		go func() {
			mq.queue.Range(func(key, value interface{}) bool {
				logger.Info("From queue")
				logger.Info(key, value)
				err := mq.client.Send(value)
				if err == nil {
					mq.queue.Delete(key)
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
	mq.length--
}

func (mq *memQueue) Get(k interface{}) (interface{}, bool) {
	return mq.queue.Load(k)
}

func (mq *memQueue) ForEach(f func(key interface{}, value interface{}) bool) {
	mq.queue.Range(func(key, value interface{}) bool {
		return f(key, value)
	})
}

func (mq *memQueue) String() string {
	return "Memory queue"
}

func NewQueue(c client.Client) Queue {
	q := &memQueue{
		client: c,
	}
	go q.handle()
	return q
}
