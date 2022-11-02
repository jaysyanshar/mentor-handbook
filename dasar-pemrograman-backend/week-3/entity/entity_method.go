package entity

import "time"

const (
	VehicleType_Car        = "Car"
	VehicleType_Motorcycle = "Motorcycle"
	VehicleType_Plane      = "Plane"
)

const (
	MaxYear_Car        = 5
	MaxYear_Motorcycle = 1
	MaxYear_Plane      = 10
)

// Car Methods
func (v Car) GetVehicleType() string {
	return VehicleType_Car
}

func (v Car) IsNew() bool {
	return getCurrentYear()-v.Year <= MaxYear_Car
}

// Motorcycle Methods
func (v Motorcycle) GetVehicleType() string {
	return VehicleType_Motorcycle
}

func (v Motorcycle) IsNew() bool {
	return getCurrentYear()-v.Year <= MaxYear_Motorcycle
}

// Plane Methods
func (v Plane) GetVehicleType() string {
	return VehicleType_Plane
}

func (v Plane) IsNew() bool {
	return getCurrentYear()-v.Year <= MaxYear_Plane
}

// private functions
func getCurrentYear() int {
	return time.Now().Year()
}
