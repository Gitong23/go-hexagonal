package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
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
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()

	customerService := service.NewCustomerService(customerRepository)
	// customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	log.Printf("Banking service running on port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	// config.Load()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
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
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}