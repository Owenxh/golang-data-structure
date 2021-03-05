package queue

import (
	"fmt"

	"io.vava.datastructure/array"
)

type ArrayQueue struct {
	array *array.Array
}

func New() *ArrayQueue {
	return &ArrayQueue{array.New()}
}

func NewWithCapacity(capacity int) *ArrayQueue {
	return &ArrayQueue{array.NewWithCapacity(capacity)}
}

func (q *ArrayQueue) Enqueue(e int) {
	q.array.AddLast(e)
}

func (q *ArrayQueue) Dequeue() int {
	return q.array.RemoveFirst()
}

func (q *ArrayQueue) Head() int {
	return q.array.GetFirst()
}

func (q *ArrayQueue) Size() int {
	return q.array.Size()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue) String() string {
	return fmt.Sprintf("Queue:{Head %v Tail, size:%d}", q.array.GetAll(), q.array.Size())
}
