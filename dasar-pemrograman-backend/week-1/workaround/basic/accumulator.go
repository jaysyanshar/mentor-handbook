package main

import (
	"fmt"
)

func main() {
	var numbers []int = Input()
	total := Sum(numbers)
	fmt.Print("Total = ", total)
}

func Input() []int {
	var n int
	var numbers []int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Print("Number ", i, ": ")
		fmt.Scan(&m)
		numbers = append(numbers, m)
	}
	return numbers
}

func Sum(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}
