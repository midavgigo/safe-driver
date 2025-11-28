package dbentry

import (
	"app/utils"
	"fmt"
	"log"
	"strconv"
)

type Status int

const (
	PENDGING                 Status = 0
	SEARCHING                Status = 1
	DRIVER_ASSIGNED          Status = 2
	WAITING_FOR_CONFIRMATION Status = 3
	IN_PROGRESS              Status = 4
	COMPLETED                Status = 5
)

func GetOrderStatusName(dbman DBManager, id int) (string, error) {
	rows, err := dbman.Query("SELECT GetOrderStatusName($1);", id)
	if err != nil {
		return "", utils.ReasonableError{
			Reason:  err,
			Message: "Error in getting status of order with id=" + strconv.Itoa(id),
		}
	}
	var ret string
	if rows.Next() {
		err := rows.Scan(&ret)
		if err != nil {
			return "", utils.ReasonableError{
				Reason:  err,
				Message: "Error in scaning rows",
			}
		}
		log.Println("Get order status", ret)
	} else {
		return "", utils.ReasonableError{
			Reason:  err,
			Message: "Can't get status of order. Look db logs",
		}
	}
	return ret, nil
}

func SetOrderStatus(dbman DBManager, id int, status Status) error {
	res, err := dbman.Exec("UPDATE Orders SET CurrentStatus = 1 WHERE Id = $1;", id)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in setting '" + fmt.Sprintf("%d", status) + "' status",
		}
	}
	n, err := res.RowsAffected()
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in getting number of affected rows",
		}
	}
	if n < 1 {
		return utils.ReasonableError{
			Reason:  nil,
			Message: "Status of order with id=" + strconv.Itoa(id) + " not changed, check id",
		}
	}
	return nil
}

func MakeOrder(
	dbman DBManager,
	address_from string,
	address_to string,
	tariff string,
	passenger_id string,
	//add services and comment
) (int, error) {
	rows, err := dbman.Query("SELECT MakeOrder($1, $2, $3, $4);",
		address_from,
		address_to,
		tariff,
		passenger_id,
	)
	if err != nil {
		return 0, utils.ReasonableError{
			Reason:  err,
			Message: "Error in calling func MakeOrder",
		}
	}
	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, utils.ReasonableError{
				Reason:  err,
				Message: "Error in scanning rows",
			}
		}
		log.Println("New order with id", id)
	} else {
		return 0, utils.ReasonableError{
			Reason:  nil,
			Message: "Can't get id of new order",
		}
	}
	return id, nil
}

func CancelOrder(dbman DBManager, id int) error {
	_, err := dbman.Exec("CALL CancelOrder($1);", id)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in canceling order",
		}
	}
	return nil
}
