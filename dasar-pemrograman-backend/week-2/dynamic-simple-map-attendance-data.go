package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Present, Absent, Sick, Present, Present, Present"
	arr := strings.Split(data, ", ")
	mappedData := make(map[string]int)
	CountAttendance(arr, mappedData)
	fmt.Println(mappedData)
}

func CountAttendance(arr []string, mappedData map[string]int) {
	for _, v := range arr {
		mappedData[v]++
	}
}
