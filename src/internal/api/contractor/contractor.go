package contractor

import (
	"app/internal/api"
	"log"
)

func DriverStatus(model DriverStatusModel) {
	log.Println("Get model", model)
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
