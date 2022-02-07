package main

// Gets the power series of integer 'a'
// and returns tuple of square of 'a' and cube of 'a'
func TupleOverview(a int) (int, int, error) {
	return a * a, a * a * a, nil
}
