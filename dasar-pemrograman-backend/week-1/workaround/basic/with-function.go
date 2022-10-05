package main

import "fmt"

func main() {
	var n, total int
	var arr []int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Print("Input number ", i, ": ")
		fmt.Scan(&m)
		arr = append(arr, m)
	}
	total = sum(arr)
	fmt.Print("Total = ", total)
}

func sum(arr []int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}
