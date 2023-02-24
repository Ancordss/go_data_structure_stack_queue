package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generar un millón de números aleatorios
	var numbers [1000000]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(20000001) - 10000000
	}

	// Insertar los números en una pila y extraerlos
	start := time.Now()
	stack := arraylist.New()
	for _, num := range numbers {
		stack.Add(num)
	}
	for i := 0; i < len(numbers); i++ {
		stack.Remove(stack.Size() - 1)
	}
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de inserción y extracción en pila: %s\n", elapsed)

	// Insertar los números en una cola y extraerlos
	start = time.Now()
	queue := arrayqueue.New()
	for _, num := range numbers {
		queue.EnQueue(num)
	}
	for i := 0; i < len(numbers); i++ {
		queue.DeQueue()
	}
	elapsed = time.Since(start)
	fmt.Printf("Tiempo de inserción y extracción en cola: %s\n", elapsed)
}
