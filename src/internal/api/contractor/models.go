package contractor

type DriverStatusModel struct {
	IsAvailable     bool
	CurrentLocation map[string]float64
}

type NewOrderStatusModel struct {
	Status string
}
