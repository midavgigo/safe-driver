package customer

import (
	"app/internal/api"
	"app/internal/api/dbentry"
	db "app/internal/api/dbentry"
	"app/utils"
	"encoding/json"
	"log"
	"strconv"
)

func print_model(text string, model any) {
	str, _ := json.Marshal(model)
	log.Println(text, string(str))
}

func MakeOrder(model MakeOrderModel, dbman db.DBManager) (int, error) {
	print_model("Get model", model)
	id, err := dbentry.MakeOrder(
		dbman,
		model.AddressFrom,
		model.AddressTo,
		model.Tariff,
		model.PassengerId,
	)
	if err != nil {
		return 0, utils.ReasonableError{
			Reason:  err,
			Message: "Error in making order with current model",
		}
	}
	err = db.SetOrderStatus(dbman, id, db.SEARCHING)
	if err != nil {
		return 0, utils.ReasonableError{
			Reason:  err,
			Message: "Error in setting status of order " + strconv.Itoa(id),
		}
	}
	return id, nil
}

func StatusOrder(request api.OrderRequest, dbman db.DBManager) (string, error) {
	print_model("Get request", request)
	order_id, err := strconv.ParseInt(request.OrderId, 10, 64)
	if err != nil {
		return "", utils.ReasonableError{
			Reason:  err,
			Message: "Error in parsing order_id to number",
		}
	}
	str, err := dbentry.GetOrderStatusName(dbman, int(order_id))
	if err != nil {
		return "", utils.ReasonableError{
			Reason:  err,
			Message: "Error in getting status of order " + request.OrderId,
		}
	}
	return str, nil
}

func CancelOrder(request api.OrderRequest, dbman dbentry.DBManager) error {
	print_model("Get request", request)
	order_id, err := strconv.ParseInt(request.OrderId, 10, 64)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in parsing order_id to number",
		}
	}
	return dbentry.CancelOrder(dbman, int(order_id))
}
