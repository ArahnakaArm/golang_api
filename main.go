package main

import (
	"fmt"
	"hello/handler"
	"hello/repository"
	"hello/service"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()
	customerRepository := repository.NewCustomerRepositoryDB(db)

	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepository
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepository)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	log.Printf("Banking service started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict

}

func initDatabase() *sqlx.DB {
	dns := fmt.Sprintf("%v:%v@/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"))

	db, err := sqlx.Open(fmt.Sprintf("%v", viper.GetString("db.driver")), dns)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
