package main

import (
	"fmt"
	"vehicle_detector/io"
	"vehicle_detector/service"
)

func main() {
	brand, model, year, err := io.Input()
	if err != nil {
		fmt.Println(err)
		return
	}
	vehicle, err := service.PopulateVehicle(brand, model, year)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Output(vehicle.GetVehicleType(), vehicle.IsNew())
}
