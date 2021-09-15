package main

import (
	"hello/handler"
	"hello/repository"
	"hello/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:12345678@/bank?parseTime=true")

	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)

	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepository

	customerService := service.NewCustomerService(customerRepositoryMock)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":3000", router)

}
