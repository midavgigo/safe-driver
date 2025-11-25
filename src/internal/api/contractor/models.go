package contractor

type DriverStatusModel struct {
	is_available     bool
	current_location []float32
}

type NewOrderStatusModel struct {
	status string
}
