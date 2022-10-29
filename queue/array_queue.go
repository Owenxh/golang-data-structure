package queue

import (
	"fmt"

	"io.vava.datastructure/array"
)

type ArrayQueue[E any] struct {
	array *array.Array[E]
}

func NewArrayQueue[E any]() *ArrayQueue[E] {
	return &ArrayQueue[E]{array.New[E]()}
}

func NewArrayQueueWithCapacity[E any](capacity int) *ArrayQueue[E] {
	return &ArrayQueue[E]{array.NewWithCapacity[E](capacity)}
}

func (q *ArrayQueue[E]) Enqueue(e E) {
	q.array.AddLast(e)
}

func (q *ArrayQueue[E]) Dequeue() E {
	return q.array.RemoveFirst()
}

func (q *ArrayQueue[E]) Head() E {
	return q.array.GetFirst()
}

func (q *ArrayQueue[E]) Size() int {
	return q.array.Size()
}

func (q *ArrayQueue[E]) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue[E]) String() string {
	return fmt.Sprintf("Queue:{Head %v Tail, size:%d}", q.array.GetAll(), q.array.Size())
}
