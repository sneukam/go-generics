/*
Practicing learning golang.
Going through the generics tutorial on go.dev.
*/

package main

import (
	"fmt"
)

// The Number interface will be used as a "type constraint"
type Number interface {
	int64 | float64
}

func main() {
	// initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// initialize a map for float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Here we used standard (non-generic) functions
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	// Below we use the generic function
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// Below we use the generic function
	fmt.Printf("Generic w/ Interface Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// Return the sum of the integers in map m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Return the sum of floats in map m
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Sum floats or ints using generics - a generic function that handles both types.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Sum the numbers using the interface Number for int64 or float64
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
