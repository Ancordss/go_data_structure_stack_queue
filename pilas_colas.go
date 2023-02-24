package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

const (
	numElements = 1000000
	rangeMin    = -10000000
	rangeMax    = 10000000
)

func generateRandomNumbers() []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		numbers[i] = rand.Intn(rangeMax-rangeMin+1) + rangeMin
	}
	return numbers
}

func insertIntoStack(numbers []int) *arraylist.List {
	start := time.Now()
	stack := arraylist.New()
	for _, num := range numbers {
		stack.Add(num)
	}
	elapsed := time.Since(start)
	fmt.Printf("tiempo para insertar %d elementos en una fila: %v\n", numElements, elapsed)
	return stack
}

func extractFromStack(stack *arraylist.List) {
	start := time.Now()
	for !stack.Empty() {
		stack.Remove(stack.Size() - 1)
	}
	elapsed := time.Since(start)
	fmt.Printf("tiempo para extraer %d elementos de una fila: %v\n", numElements, elapsed)
}

func insertIntoQueue(numbers []int) *arrayqueue.Queue {
	start := time.Now()
	queue := arrayqueue.New()
	for _, num := range numbers {
		queue.Enqueue(num)
	}
	elapsed := time.Since(start)
	fmt.Printf("tiempo para insertar %d elementos en una cola: %v\n", numElements, elapsed)
	return queue
}

func extractFromQueue(queue *arrayqueue.Queue) {
	start := time.Now()
	for queue.Size() > 0 {
		queue.Dequeue()
		//fmt.Printf("%d ", num)
	}
	elapsed := time.Since(start)
	fmt.Printf("tiempo para extraer %d elementos de uan cola: %v\n", numElements, elapsed)
}

func main() {
	numbers := generateRandomNumbers()

	stack := insertIntoStack(numbers)
	extractFromStack(stack)

	queue := insertIntoQueue(numbers)
	extractFromQueue(queue)
}
