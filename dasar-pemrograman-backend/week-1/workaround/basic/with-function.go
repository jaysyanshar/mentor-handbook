package main

import "fmt"

func main() {
	var n int
	var arr []int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Print("Input number ", i, ": ")
		fmt.Scan(&m)
		arr = append(arr, m)
	}
	sum := sum(arr)
	odd := countOdd(arr)
	fmt.Println("Sum =", sum)
	fmt.Println("Count Odd =", odd)
}

func sum(arr []int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}

func countOdd(arr []int) int {
	var total int
	for _, v := range arr {
		if v%2 != 0 {
			total++
		}
	}
	return total
}
