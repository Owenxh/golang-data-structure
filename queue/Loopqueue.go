package queue

import (
	"fmt"
	"strconv"
	"strings"
)

type LoopQueue struct {
	array []int
	size  int
	head  int
	tail  int
}

func NewLoopQueue() *LoopQueue {
	return NewLoopQueueWithCapacity(8)
}

func NewLoopQueueWithCapacity(c int) *LoopQueue {
	return &LoopQueue{
		array: make([]int, c, c),
		size:  0,
		head:  0,
		tail:  0,
	}
}

func (q *LoopQueue) capacity() int {
	return len(q.array)
}

func (q *LoopQueue) resize(c int) {
	newArr := make([]int, c, c)

	for i := 0; i < q.size; i++ {
		newArr[i] = q.array[(i+q.head)%q.capacity()]
	}

	q.array = newArr
	q.head, q.tail = 0, q.size
}

func (q *LoopQueue) Enqueue(e int) {
	if q.size == q.capacity() {
		q.resize((q.capacity()) * 2)
	}
	q.array[q.tail] = e
	q.tail = (q.tail + 1) % q.capacity()
	q.size++
}

func (q *LoopQueue) Dequeue() int {
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

func (q *LoopQueue) Head() int {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}

	return q.array[q.head]
}

func (q *LoopQueue) Size() int {
	return q.size
}

func (q *LoopQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LoopQueue) String() string {
	builder := strings.Builder{}
	//index := q.head
	builder.WriteString("[")
	//for ; index != q.tail; index = (index + 1) % q.capacity() {
	//	v := q.array[index%q.capacity()]
	//	builder.WriteString(strconv.Itoa(v))
	//
	//	if (index+1)%q.capacity() != q.tail {
	//		builder.WriteString(",")
	//	}
	//}
	for i := 0; i < q.size; i++ {
		v := q.array[(i+q.head)%q.capacity()]
		builder.WriteString(strconv.Itoa(v))

		if (i+1)%q.capacity() != q.tail {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")

	return fmt.Sprintf("Queue:{Head %v Tail, size:%d, head:%d, tail:%d, capacity:%d}",
		builder.String(), q.size, q.head, q.tail, q.capacity())
}
