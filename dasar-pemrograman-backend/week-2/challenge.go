package main

import (
	"fmt"
	"strings"
)

const (
	PRESENT = "Present"
	ABSENT  = "Absent"
	SICK    = "Sick"
)

func main() {
	input := "Present, Present, Absent, Present, Sick"
	arr := InputToArray(input)
	result := MapAttendance(arr)
	fmt.Println(result)
}

func InputToArray(input string) []string {
	return strings.Split(input, ", ")
}
func MapAttendance(arr []string) map[string]int {
	attendances := make(map[string]int)
	attendances[PRESENT] = 0
	attendances[ABSENT] = 0
	attendances[SICK] = 0
	for _, v := range arr {
		if !isValidAttendance(v) {
			continue
		}
		attendances[v]++
	}
	return attendances
}

func isValidAttendance(val string) bool {
	return val == PRESENT || val == ABSENT || val == SICK
}
