package main

import (
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
}
