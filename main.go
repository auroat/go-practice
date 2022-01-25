package main

import (
	"container/heap"
	"fmt"
)

func main() {

	// list-overview
	ListOverview()

	// tuple-overview
	var square int
	var cube int
	square, cube, _ = TupleOverview(3)
	fmt.Println("Square result:", square, "\nCube result:", cube)

	// heap-overview
	var intHeap *IntegerHeap = &IntegerHeap{-7, 1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
