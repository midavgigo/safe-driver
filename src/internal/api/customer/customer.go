package customer

import (
	"app/internal/api"
	"log"
)

func MakeOrder(model MakeOrderModel) {
	log.Println("Get model", model)
}

func StatusOrder(request api.OrderRequest) {
	log.Println("Get request", request)
}

func CancelOrder(request api.OrderRequest) {
	log.Println("Get request", request)
}
