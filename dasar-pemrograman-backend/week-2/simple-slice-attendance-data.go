package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Person_1, Present, Absent, Sick, Present, Present, Present, Person_2, Absent, Sick, Present"
	arr := strings.Split(data, ", ")
	slices := getAttendances(arr)
	fmt.Println(slices)
}

func getAttendances(data []string) [][]string {
	slices := make([][]string, 0)
	i := -1
	for _, v := range data {
		if v == "Present" || v == "Absent" || v == "Sick" {
			slices[i] = append(slices[i], v)
		} else {
			slices = append(slices, make([]string, 0))
			i++
		}
	}
	return slices
}
