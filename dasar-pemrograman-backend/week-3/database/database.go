package database

import (
	"vehicle_detector/entity"
)

type VehicleData struct {
	Type    string
	Vehicle entity.Vehicle
}

type Database interface {
	Init()
	Load() []VehicleData
	IsActive() bool
}
