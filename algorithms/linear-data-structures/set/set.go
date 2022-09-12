package main

import (
	"fmt"
)

type Set struct {
	integerMap map[int]bool
}

// Creates the map of integer-key and bool-value
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// Adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// Deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// Checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

// Returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int

	for value = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
	fmt.Println(set.ContainsElement(3))
}
