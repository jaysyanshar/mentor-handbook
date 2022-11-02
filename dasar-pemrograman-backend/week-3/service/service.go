package service

import (
	"errors"
	"vehicle_detector/database"
	"vehicle_detector/entity"
)

var db database.Database = &database.ListDatabase{}

func PopulateVehicle(brand, model string, year int) (entity.Vehicle, error) {
	if !db.IsActive() {
		db.Init()
	}
	for _, v := range db.Load() {
		if v.Type == entity.VehicleType_Car {
			car := v.Vehicle.(entity.Car)
			if car.Brand == brand && car.Model == model {
				return v.Vehicle, nil
			}
		} else if v.Type == entity.VehicleType_Motorcycle {
			motorcycle := v.Vehicle.(entity.Motorcycle)
			if motorcycle.Brand == brand && motorcycle.Model == model {
				return v.Vehicle, nil
			}
		} else if v.Type == entity.VehicleType_Plane {
			plane := v.Vehicle.(entity.Plane)
			if plane.Brand == brand && plane.Model == model {
				return v.Vehicle, nil
			}
		}
	}
	return nil, errors.New("vehicle not exists")
}
