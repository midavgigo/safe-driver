package contractor

import (
	"app/internal/api"
	"app/internal/api/dbentry"
	"encoding/json"
	"log"
)

func print_model(text string, model any) {
	str, _ := json.Marshal(model)
	log.Println(text, string(str))
}

func DriverStatus(model DriverStatusModel, dbman dbentry.DBManager) error {
	print_model("Get model", model)
	return dbentry.SetDriverStatus(dbman, 0, model.IsAvailable, model.CurrentLocation["lat"], model.CurrentLocation["lng"])
}

func AcceptOrder(request api.OrderRequest) {
	log.Println("Get request", request)
}

func DriverArrived(request api.OrderRequest) {
	log.Println("Get request", request)
}

func ChangeOrderStatus(request api.OrderRequest, model NewOrderStatusModel) {
	log.Println("Get model", model)
	log.Println("Get request", request)
}
