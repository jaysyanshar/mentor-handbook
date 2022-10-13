package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Present, Absent, Sick, Present, Present, Present"
	arr := strings.Split(data, ", ")
	mappedData := make(map[string]int, 3)
	mappedData["Present"] = 0
	mappedData["Absent"] = 0
	mappedData["Sick"] = 0
	CountAttendance(arr, mappedData)
	fmt.Println(mappedData)
}

func CountAttendance(arr []string, mappedData map[string]int) {
	for _, v := range arr {
		if v == "Present" {
			mappedData["Present"]++
		} else if v == "Absent" {
			mappedData["Absent"]++
		} else if v == "Sick" {
			mappedData["Sick"]++
		}
	}
}
