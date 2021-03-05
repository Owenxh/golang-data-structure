package queue

type IntQueue interface {
	Enqueue(int)

	Dequeue() int

	Head() int

	Size() int

	IsEmpty() bool
}
