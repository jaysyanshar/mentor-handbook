package entity

type Vehicle interface {
	GetVehicleType() string
	IsNew() bool
}

type Car struct {
	Brand string
	Model string
	Year  int
}

type Motorcycle struct {
	Brand string
	Model string
	Year  int
}

type Plane struct {
	Brand string
	Model string
	Year  int
}
