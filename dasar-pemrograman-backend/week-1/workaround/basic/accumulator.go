package main

import "fmt"

func main() {
	var n, total int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Print("Input number ", i, ": ")
		fmt.Scan(&m)
		total += m
	}
	fmt.Print("Total = ", total)
}
