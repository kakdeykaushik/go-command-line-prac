package models

type Queue interface {
	Enqueue(any)
	Dequeue()
}
