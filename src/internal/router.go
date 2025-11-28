package internal

import (
	"app/internal/api"
	"app/internal/api/contractor"
	"app/internal/api/customer"
	"app/internal/api/dbentry"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func unsafe_options(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.WriteHeader(http.StatusOK)
}
func ApiOrder(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()

		var decoded map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Println("Error in decoding body /api/order")
			return
		}
		var services []string
		for _, v := range decoded["selected_services"].([]interface{}) {
			services = append(services, v.(string))
		}

		model := customer.MakeOrderModel{
			PassengerId:      decoded["passenger_id"].(string),
			AddressFrom:      decoded["address_from"].(string),
			AddressTo:        decoded["address_to"].(string),
			Tariff:           decoded["tariff"].(string),
			SelectedServices: services,
			Comment:          decoded["comment"].(string),
		}
		id, err := customer.MakeOrder(model, dbman)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "New order with id: %d\n", id)
	case http.MethodOptions:
		unsafe_options(w, r)
	}
}

func ApiOrderStatus(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order/{order_id}. Method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		status, err := customer.StatusOrder(
			api.OrderRequest{
				OrderId: strings.Split(r.URL.Path, "/")[3],
			},
			dbman,
		)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "Status of order: %s\n", status)
	}
}

func ApiOrderCancel(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order/{order_id}/cancel. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		err := customer.CancelOrder(
			api.OrderRequest{
				OrderId: strings.Split(r.URL.Path, "/")[3],
			},
			dbman,
		)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "Order successfully canceled")
	}
}

func ApiDriverStatus(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/driver/status. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()

		var decoded map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Println("Error in decoding body /api/order")
			return
		}
		current_location := make(map[string]float64)
		raw_map := decoded["current_location"].(map[string]interface{})
		for k, v := range raw_map {
			current_location[k], _ = strconv.ParseFloat(v.(string), 32)
		}
		err = contractor.DriverStatus(
			contractor.DriverStatusModel{
				IsAvailable:     decoded["is_available"].(bool),
				CurrentLocation: current_location,
			},
			dbman,
		)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintln(w, "Driver status updated")
	case http.MethodOptions:
		unsafe_options(w, r)
	}
}

func ApiOrderAccept(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order/{order_id}/accept. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		contractor.AcceptOrder(
			api.OrderRequest{
				OrderId: strings.Split(r.URL.Path, "/")[3],
			},
		)
	}
}

func ApiOrderArrived(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order/{order_id}/arrived. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		contractor.DriverArrived(
			api.OrderRequest{
				OrderId: strings.Split(r.URL.Path, "/")[3],
			},
		)
	}
}

func ApiOrderNewStatus(w http.ResponseWriter, r *http.Request, dbman dbentry.DBManager) {
	log.Println("Get request for /api/order/{order_id}/accept. Method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()

		var decoded map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Println("Error in decoding body /api/order")
			return
		}
		contractor.ChangeOrderStatus(
			api.OrderRequest{
				OrderId: strings.Split(r.URL.Path, "/")[3],
			},
			contractor.NewOrderStatusModel{
				Status: decoded["status"].(string),
			},
		)
	case http.MethodOptions:
		unsafe_options(w, r)
	}
}
