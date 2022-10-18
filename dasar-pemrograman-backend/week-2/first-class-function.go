package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// IIFE
	sum, mul := func(a []int) (int, int) {
		s := Sum(a)
		m := Multiply(a)
		return s, m
	}(arr)
	fmt.Println("sum =", sum)
	fmt.Println("mul =", mul)
}

// normal function
func Sum(arr []int) int {
	// function as variable
	add := func(a, b int) int {
		return a + b
	}
	total := 0
	Loop(arr, func(x int) {
		// closure
		total = add(total, x)
	})
	return total
}

func Multiply(arr []int) int {
	total := 1
	Loop(arr, func(x int) {
		mul(&total, x)
	})
	return total
}

// pointer
func mul(current *int, x int) {
	*current = *current * x
}

// function as parameter
func Loop(arr []int, apply func(x int)) {
	for _, v := range arr {
		apply(v)
	}
}
