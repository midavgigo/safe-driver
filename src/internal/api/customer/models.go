package customer

type MakeOrderModel struct {
	passenger_id      string
	address_from      string
	address_to        string
	tariff            string
	selected_services []string
	comment           string
}
