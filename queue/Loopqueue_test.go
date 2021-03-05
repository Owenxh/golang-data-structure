package queue_test

import (
	"fmt"
	"testing"

	"io.vava.datastructure/queue"
)

func TestLoopQueue_Enqueue(t *testing.T) {
	queue := queue.NewLoopQueueWithCapacity(8)

	n := 16
	for i := 0; i < n; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}
	fmt.Printf("Head = %v, empty = %v\n", queue.Head(), queue.IsEmpty())

	for i := 0; i < n; i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

	for i := 0; i < n; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}
	fmt.Printf("Head = %v, empty = %v\n", queue.Head(), queue.IsEmpty())

	for i := 0; i < n; i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}
}
