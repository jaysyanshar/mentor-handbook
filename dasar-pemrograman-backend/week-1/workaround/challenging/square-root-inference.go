package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)
	sqRoot := SquareRoot(n)
	fmt.Print("Square Root = ", sqRoot)
}

func SquareRoot(n int) int {
	x := n
	for a := 1; a <= n; a++ {
		x = x - ((x*x - a) / (2 * x))
	}
	return x
}
