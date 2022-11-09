package queue

import (
	"fmt"
	"strings"
)

type LoopQueue[E any] struct {
	array []E
	size  int
	head  int
	tail  int
}

func NewLoopQueue[E any]() *LoopQueue[E] {
	return NewLoopQueueWithCapacity[E](8)
}

func NewLoopQueueWithCapacity[E any](c int) *LoopQueue[E] {
	return &LoopQueue[E]{
		array: make([]E, c),
		size:  0,
		head:  0,
		tail:  0,
	}
}

func (q *LoopQueue[E]) capacity() int {
	return len(q.array)
}

func (q *LoopQueue[E]) resize(c int) {
	newArr := make([]E, c)

	for i := 0; i < q.size; i++ {
		newArr[i] = q.array[(i+q.head)%q.capacity()]
	}

	q.array = newArr
	q.head, q.tail = 0, q.size
}

func (q *LoopQueue[E]) Enqueue(e E) {
	if q.size == q.capacity() {
		q.resize((q.capacity()) * 2)
	}
	q.array[q.tail] = e
	q.tail = (q.tail + 1) % q.capacity()
	q.size++
}

func (q *LoopQueue[E]) Dequeue() E {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}

	res := q.array[q.head]

	q.head = (q.head + 1) % q.capacity()
	q.size--

	if q.size > 0 && q.size == (q.capacity())/4 {
		q.resize((q.capacity()) / 2)
	}

	return res
}

func (q *LoopQueue[E]) Head() E {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}

	return q.array[q.head]
}

func (q *LoopQueue[E]) Size() int {
	return q.size
}

func (q *LoopQueue[E]) IsEmpty() bool {
	return q.size == 0
}

func (q *LoopQueue[E]) String() string {
	builder := strings.Builder{}
	// index := q.head
	builder.WriteString("[")
	// for ; index != q.tail; index = (index + 1) % q.capacity() {
	//	v := q.array[index%q.capacity()]
	//	builder.WriteString(strconv.Itoa(v))
	//
	//	if (index+1)%q.capacity() != q.tail {
	//		builder.WriteString(",")
	//	}
	//}
	for i := 0; i < q.size; i++ {
		v := q.array[(i+q.head)%q.capacity()]
		builder.WriteString(fmt.Sprintf("%v", v))

		if (i+1)%q.capacity() != q.tail {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")

	return fmt.Sprintf("Queue:{Head %v Tail, size:%d, head:%d, tail:%d, capacity:%d}",
		builder.String(), q.size, q.head, q.tail, q.capacity())
}
