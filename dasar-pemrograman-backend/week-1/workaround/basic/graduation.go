package main

import "fmt"

func checkGraduation(score int, absen int) string {
	var predikat string
	if score >= 70 && absen < 5 {
		predikat = "lulus"
	} else {
		predikat = "tidak lulus"
	}
	return predikat
}

func main() {
	fmt.Println(checkGraduation(60, 4))
}
