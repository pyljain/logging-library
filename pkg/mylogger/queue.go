package mylogger

import (
	"sync"
	"time"
)

type queue struct {
	maxSize              int
	defaultFlushInterval time.Duration
	queue                []*queueElement
	flushFn              func([]*queueElement)
	sync.Mutex
}

type queueElement struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Elements  map[string]interface{} `json:"elements"`
}

func newQueue(maxSize int, defaulFlushInterval time.Duration, flushFn func([]*queueElement)) *queue {
	return &queue{
		maxSize:              maxSize,
		defaultFlushInterval: defaulFlushInterval,
		flushFn:              flushFn,
	}
}

func (q *queue) add(qe *queueElement) {
	q.Lock()
	q.queue = append(q.queue, qe)
	q.Unlock()
	if len(q.queue) == q.maxSize {
		go q.flush()
	}
}

func (q *queue) flush() {
	q.Lock()
	defer q.Unlock()
	go q.flushFn(q.queue[:])
	q.queue = []*queueElement{}
}

func (q *queue) run() {
	for {
		time.Sleep(q.defaultFlushInterval)
		q.flush()
	}
}
