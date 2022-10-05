package main

import "fmt"

func main() {
	var gender, name, title string
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Gender: ")
	fmt.Scan(&gender)
	if gender == "male" {
		title = "Mr"
	} else if gender == "female" {
		title = "Mrs"
	}
	fmt.Print("Hi ", title, " ", name)
}
