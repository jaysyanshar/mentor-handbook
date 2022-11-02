package database

import "vehicle_detector/entity"

type ListDatabase struct {
	data []VehicleData
}

func (d *ListDatabase) Init() {
	d.data = make([]VehicleData, 0)
	d.data = append(d.data,
		VehicleData{
			Type: entity.VehicleType_Car,
			Vehicle: entity.Car{
				Brand: "Toyota",
				Model: "Yaris",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Car,
			Vehicle: entity.Car{
				Brand: "Honda",
				Model: "Jazz",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Car,
			Vehicle: entity.Car{
				Brand: "Toyota",
				Model: "Rush",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Car,
			Vehicle: entity.Car{
				Brand: "Mitsubishi",
				Model: "Xpander",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Motorcycle,
			Vehicle: entity.Motorcycle{
				Brand: "Yamaha",
				Model: "XSR",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Motorcycle,
			Vehicle: entity.Motorcycle{
				Brand: "Honda",
				Model: "Vario",
			},
		},
		VehicleData{
			Type: entity.VehicleType_Plane,
			Vehicle: entity.Plane{
				Brand: "Boeing",
				Model: "737",
			},
		},
	)
}

func (d *ListDatabase) Load() []VehicleData {
	return d.data
}

func (d *ListDatabase) IsActive() bool {
	return len(d.data) > 0
}
