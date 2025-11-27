package internal

import (
	"log"
	"net/http"
)

const port string = "8080"
const host string = ""

func StartServer() *http.Server {
	server := &http.Server{
		Addr: host + ":" + port,
	}
	HandleFuncs()
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
	return server
}

func HandleFuncs() {
	http.HandleFunc("/api/order", ApiOrder)
	http.HandleFunc("/api/order/{order_id}", ApiOrderStatus)
	http.HandleFunc("/api/order/{order_id}/cancel", ApiOrderCancel)
	http.HandleFunc("/api/driver/status", ApiDriverStatus)
	http.HandleFunc("/api/order/{order_id}/accept", ApiOrderAccept)
	http.HandleFunc("/api/order/{order_id}/arrived", ApiOrderArrived)
	http.HandleFunc("/api/order/{order_id}/status", ApiOrderNewStatus)
}
