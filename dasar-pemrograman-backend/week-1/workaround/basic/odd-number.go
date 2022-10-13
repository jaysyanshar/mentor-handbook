package main

import "fmt"

func main() {
	var n int
	var numbers []int
	fmt.Print("Number of inputs: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var input int
		fmt.Print("Number ", i, ": ")
		fmt.Scan(&input)
		numbers = append(numbers, input)
	}
	odd := CountOdds(numbers)
	fmt.Println("Odd:", odd)
}

func CountOdds(numbers []int) int {
	var total int
	for _, v := range numbers {
		if v%2 != 0 {
			total++
		}
	}
	return total
}
