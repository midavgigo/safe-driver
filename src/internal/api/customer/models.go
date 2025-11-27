package customer

type MakeOrderModel struct {
	PassengerId      string
	AddressFrom      string
	AddressTo        string
	Tariff           string
	SelectedServices []string
	Comment          string
}
