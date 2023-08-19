package queue

import (
	"errors"
)

// fixed size Queue implementation using buffered channel
// todo: make queue item generic
type queue struct {
	item chan string
}

func NewQueue(size int) (*queue, error) {
	if size < 1 {
		return nil, errors.New("invalid size")
	}

	return &queue{item: make(chan string, size)}, nil
}

func (q *queue) Enqueue(val string) {
	if q.isFull() {
		q.Dequeue()
	}

	q.item <- val
}

func (q *queue) Dequeue() (string, error) {
	if !q.isEmpty() {
		return <-q.item, nil
	}

	return "", errors.New("empty queue")
}

func (q queue) isFull() bool {
	return len(q.item) == cap(q.item)
}

func (q queue) isEmpty() bool {
	return len(q.item) == 0
}
