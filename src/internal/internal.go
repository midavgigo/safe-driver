package internal

import (
	"app/internal/api/dbentry"
	"app/utils"
	"log"
	"net/http"
)

const port string = "8080"
const host string = ""

func StartServer() (*http.Server, error) {
	server := &http.Server{
		Addr: host + ":" + port,
	}
	dbman, err := dbentry.New()
	if err != nil {
		return nil, utils.ReasonableError{
			Reason:  err,
			Message: "Error in starting server",
		}
	}
	HandleFuncs(dbman)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Println("Server was closed")
			default:
				log.Println("Error in server:", err)
			}
		}
	}()
	return server, err
}

func HandleFuncs(dbman dbentry.DBManager) {
	http.HandleFunc("/api/order", func(w http.ResponseWriter, r *http.Request) {
		ApiOrder(w, r, dbman)
	})
	http.HandleFunc("/api/order/{order_id}", func(w http.ResponseWriter, r *http.Request) {
		ApiOrderStatus(w, r, dbman)
	})
	http.HandleFunc("/api/order/{order_id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		ApiOrderCancel(w, r, dbman)
	})
	http.HandleFunc("/api/driver/status", func(w http.ResponseWriter, r *http.Request) {
		ApiDriverStatus(w, r, dbman)
	})
	http.HandleFunc("/api/order/{order_id}/accept", func(w http.ResponseWriter, r *http.Request) {
		ApiOrderAccept(w, r, dbman)
	})
	http.HandleFunc("/api/order/{order_id}/arrived", func(w http.ResponseWriter, r *http.Request) {
		ApiOrderArrived(w, r, dbman)
	})
	http.HandleFunc("/api/order/{order_id}/status", func(w http.ResponseWriter, r *http.Request) {
		ApiOrderNewStatus(w, r, dbman)
	})
}
