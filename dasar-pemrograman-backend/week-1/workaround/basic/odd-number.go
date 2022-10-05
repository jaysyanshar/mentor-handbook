package main

import "fmt"

func main() {
	var n int
	var odds []int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Print("Input number ", i, ": ")
		fmt.Scan(&m)
		if m%2 != 0 {
			odds = append(odds, m)
		}
	}
	fmt.Print("Odd numbers: ", odds)
}
