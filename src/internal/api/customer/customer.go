package customer

import (
	"app/internal/api"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func MakeOrder(model MakeOrderModel) {
	log.Println("Get model", model)
	connStr := "postgresql://doc:password@db:5432/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error in opening db", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT MakeOrder($1, $2, $3, $4);",
		model.AddressFrom,
		model.AddressTo,
		model.Tariff,
		model.PassengerId,
	)
	if err != nil {
		log.Println("Error in query db", err)
		return
	}
	if rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Println("Error in scaning rows", err)
			return
		}
		log.Println("New order with id", id)
	}
}

func StatusOrder(request api.OrderRequest) {
	log.Println("Get request", request)
}

func CancelOrder(request api.OrderRequest) {
	log.Println("Get request", request)
}
