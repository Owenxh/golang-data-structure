package queue

type Interface[E any] interface {
	Enqueue(E)

	Dequeue() E

	Head() E

	Size() int

	IsEmpty() bool
}
