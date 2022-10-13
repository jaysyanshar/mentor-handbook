package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Present, Absent, Sick, Present, Present, Present"
	arr := getArr(data)
	fmt.Println(arr)
	replaceAbsent(arr)
	fmt.Println(arr)
}

func getArr(s string) []string {
	s = strings.ReplaceAll(s, " ", "")
	return strings.Split(s, ",")
}

func replaceAbsent(a []string) {
	for i := 0; i < len(a); i++ {
		if a[i] == "Absent" {
			a[i] = "Present"
		}
	}
}
