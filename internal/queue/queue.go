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
	sync.RWMutex
	queue    sync.Map
	length   int
	interval time.Duration
	client   client.Client
}

func (mq *memQueue) handle() int {
	for {
		<-time.After(mq.interval)
		logger.Debug("Queue [status]Total: ", mq.Length())
		if mq.Length() > 0 {
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
}

func (mq *memQueue) Length() int {
	mq.RLock()
	defer mq.RUnlock()
	return mq.length
}

func (mq *memQueue) Push(data interface{}) {
	mq.Lock()
	mq.length++
	mq.Unlock()
	mq.queue.Store(mq.length, data)
}

func (mq *memQueue) Pop(key interface{}) {
	if mq.Length() == 0 {
		return
	}
	_, ok := mq.Get(key)
	if ok {
		mq.queue.Delete(key)
		mq.Lock()
		mq.length--
		mq.Unlock()
	}
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
