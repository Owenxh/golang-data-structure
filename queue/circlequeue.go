package queue

import (
	"fmt"
	"strconv"
	"strings"
)

type CircleQueue struct {
	array    []int
	capacity int
	size     int
	head     int
	tail     int
}

func NewCircleQueue() *CircleQueue {
	return NewCircleQueueWithCapacity(11)
}

func NewCircleQueueWithCapacity(c int) *CircleQueue {
	realCap := c + 1
	return &CircleQueue{
		array:    make([]int, realCap, realCap),
		capacity: realCap,
		size:     0,
		head:     0,
		tail:     0,
	}
}

func (q *CircleQueue) resize(c int) {
	newArr := make([]int, c, c)

	for i := 0; i < q.size; i++ {
		newArr[i] = q.array[(i+q.head)%q.capacity]
	}

	q.array = newArr
	q.capacity = c
	q.head, q.tail = 0, q.size
}

func (q *CircleQueue) Enqueue(e int) {
	if q.size == q.capacity-1 {
		q.resize((q.capacity-1)*2 + 1)
	}

	q.array[(q.tail)%q.capacity] = e
	q.tail = (q.tail + 1) % q.capacity
	q.size++
}

func (q *CircleQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}

	res := q.array[q.head]

	q.head = (q.head + 1) % q.capacity
	q.size--

	if q.size == (q.capacity-1)/4 {
		q.resize((q.capacity-1)/2 + 1)
	}

	return res
}

func (q *CircleQueue) Head() int {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}

	return q.array[q.head]
}

func (q *CircleQueue) Size() int {
	return q.size
}

func (q *CircleQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircleQueue) String() string {
	builder := strings.Builder{}
	index := q.head
	builder.WriteString("[")
	for ; index != q.tail; index = (index + 1) % q.capacity {
		v := q.array[index%q.capacity]
		builder.WriteString(strconv.Itoa(v))

		if (index+1)%q.capacity != q.tail {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")

	return fmt.Sprintf("Queue:{Head %v Tail, size:%d, head:%d, tail:%d, capacity:%d}",
		builder.String(), q.size, q.head, q.tail, q.capacity)
}
