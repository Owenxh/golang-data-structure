package queue_test

import (
	"fmt"
	"testing"

	"io.vava.datastructure/queue"
)

func TestCircleQueue_Enqueue(t *testing.T) {
	queue := queue.NewCircleQueueWithCapacity(10)

	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	fmt.Printf("Head element = %v\n", queue.Head())
	fmt.Printf("Is empty = %v\n", queue.IsEmpty())

	for i := 0; i < 20; i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

	fmt.Printf("Is empty = %v\n", queue.IsEmpty())
}
