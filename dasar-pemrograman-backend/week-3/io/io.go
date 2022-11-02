package io

import (
	"errors"
	"fmt"
	"time"
)

func Input() (brand, model string, year int, err error) {
	fmt.Print("Brand: ")
	fmt.Scan(&brand)
	fmt.Print("Model: ")
	fmt.Scan(&model)
	fmt.Print("Year: ")
	_, err = fmt.Scan(&year)
	if err == nil && year > time.Now().Year() {
		err = errors.New("invalid year")
	}
	return
}

func Output(vehicleType string, isNew bool) {
	var label string
	if isNew {
		label = "New"
	} else {
		label = "Old"
	}
	fmt.Println(label, vehicleType)
}
