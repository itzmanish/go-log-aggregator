package queue

import (
	"sync"
	"time"

	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
)

type Queue interface {
	Init(opts ...Option)
	Options() *Options
	Pop(key interface{})
	Push(value interface{})
	Get(key interface{}) (interface{}, bool)
	Length() int
	String() string
}

type memQueue struct {
	sync.RWMutex
	queue  sync.Map
	length int
	opts   Options
}

func NewQueue(opts ...Option) Queue {
	q := &memQueue{}
	q.Init(opts...)
	go q.handle()
	return q
}

func (mq *memQueue) Init(opts ...Option) {
	for _, o := range opts {
		o(&mq.opts)
	}
}

func (mq *memQueue) Options() *Options {
	return &mq.opts
}

func (mq *memQueue) handle() int {
	for {
		<-time.After(mq.opts.Interval)
		logger.Debug("Queue [status]Total: ", mq.Length())
		if mq.Length() > 0 {
			go func() {
				mq.queue.Range(func(key, value interface{}) bool {
					logger.Info(key, value)
					res := &codec.Packet{}
					err := mq.opts.Client.SendAndRecv(value, res)
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
	if mq.Length() >= mq.opts.MaxQueueSize {
		logger.Error("Queue is full...")
		return
	}
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
