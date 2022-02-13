package main

import (
	"fmt"
)

func main() {
	var rows int
	var cols int
	rows = 7
	cols = 9

	/*
	   The make built-in function allocates and initializes an object of type
	    slice, map, or chan (only). Like new, the first argument is a type, not a
	    value. Unlike new, make's return type is the same as the type of its
	    argument, not a pointer to it.
	*/
	var twodslices = make([][]int, rows)
	var i int
	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}
	fmt.Println(twodslices)

	var arr = []int{5, 6, 7, 8, 9}
	var slic1 = arr[:3]
	fmt.Println("slice1", slic1)

	// 	slic2 is a sub slice of arr starting from 1
	// (inclusive) to 5 (excluded)
	var slic2 = arr[1:5]
	fmt.Println("slice2", slic2)
	var slic3 = append(slic2, 12)
	fmt.Println("slice3", slic3)
}
