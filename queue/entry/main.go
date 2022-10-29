package main

import (
	"fmt"
	"time"

	"io.vava.datastructure/queue"
)

func main() {
	c := 10000
	benchmarkDequeue(queue.NewArrayQueueWithCapacity[int](c), c)
	benchmarkDequeue(queue.NewLoopQueueWithCapacity(c), c)

	c = 100000
	benchmarkDequeue(queue.NewArrayQueueWithCapacity[int](c), c)
	benchmarkDequeue(queue.NewLoopQueueWithCapacity(c), c)
}

func benchmarkDequeue(q queue.Interface[int], c int) {
	for i := 0; i < c; i++ {
		q.Enqueue(i)
	}

	start := time.Now()
	for i := 0; i < c; i++ {
		q.Dequeue()
	}
	fmt.Printf("Queue(size = %6d) Dequeue method cost %v\n", c, time.Now().Sub(start))
}
