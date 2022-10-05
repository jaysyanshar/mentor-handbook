package main

import "fmt"

func main() {
	var gender, name, title string
	var age int
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Gender: ")
	fmt.Scan(&gender)
	fmt.Print("Age: ")
	fmt.Scan(&age)
	if gender == "male" {
		if age >= 17 {
			title = "Mr"
		} else {
			title = "Boy"
		}
	} else if gender == "female" {
		if age >= 17 {
			title = "Ms"
		} else {
			title = "Girl"
		}
	}
	fmt.Print("Hi ", title, " ", name)
}
