package contractor

import (
	"app/internal/api"
	"app/internal/api/dbentry"
	"app/utils"
	"encoding/json"
	"log"
	"strconv"
)

func print_model(text string, model any) {
	str, _ := json.Marshal(model)
	log.Println(text, string(str))
}

func DriverStatus(model DriverStatusModel, dbman dbentry.DBManager) error {
	print_model("Get model", model)
	return dbentry.SetDriverStatus(dbman, 0, model.IsAvailable, model.CurrentLocation["lat"], model.CurrentLocation["lng"])
}

func AcceptOrder(request api.OrderRequest, dbman dbentry.DBManager) error {
	print_model("Get request", request)
	order_id, err := strconv.ParseInt(request.OrderId, 10, 64)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in parsing order_id to number",
		}
	}
	return dbentry.AcceptOrder(dbman, int(order_id), 0)
}

func DriverArrived(request api.OrderRequest, dbman dbentry.DBManager) error {
	print_model("Get request", request)
	order_id, err := strconv.ParseInt(request.OrderId, 10, 64)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in parsing order_id to number",
		}
	}
	return dbentry.SetOrderStatus(dbman, int(order_id), dbentry.WAITING_FOR_CONFIRMATION)
}

func ChangeOrderStatus(request api.OrderRequest, model NewOrderStatusModel, dbman dbentry.DBManager) error {
	log.Println("Get model", model)
	log.Println("Get request", request)

	order_id, err := strconv.ParseInt(request.OrderId, 10, 64)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in parsing order_id to number",
		}
	}

	status_id, err := dbentry.GetOrderStatusId(dbman, model.Status)
	if err != nil {
		return err
	}
	return dbentry.SetOrderStatus(dbman, int(order_id), status_id)
}
