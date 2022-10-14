package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Person_1, Present, Absent, Sick, Present, Present, Present, Person_2, Absent, Sick, Present"
	arr := strings.Split(data, ", ")
	attendances := getAttendances(arr)
	fmt.Println(attendances)
}

func getAttendances(arr []string) map[string]map[string]int {
	attendances := make(map[string]map[string]int)
	name := ""
	for _, v := range arr {
		if v == "Present" || v == "Absent" || v == "Sick" {
			attendances[name][v]++ // attendances["Person_1"]["Present"] += 1
		} else {
			name = v
			attendances[name] = make(map[string]int) // attendances["Person_1"] = new map()
		}
	}
	return attendances
}
