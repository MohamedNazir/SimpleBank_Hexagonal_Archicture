package api

import (
	"time"

	"github.com/MohamedNazir/SimpleBank/controller"
	"github.com/MohamedNazir/SimpleBank/domain"
	"github.com/MohamedNazir/SimpleBank/logger"
	"github.com/MohamedNazir/SimpleBank/service"
	"github.com/jmoiron/sqlx"
)

func mapUrls() {

	// wiring
	//cc := controller.CustomerController{Service: service.CustomerService(domain.NewCustomerRepositoryStub())}
	//	cc := controller.CustomerController{Service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	CustomerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	CustomerService := service.NewCustomerService(CustomerRepositoryDb)

	AccountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	AccountService := service.NewAccountService(AccountRepositoryDb)

	cc := controller.CustomerController{CustomerService}
	ac := controller.AccountController{AccountService}

	router.GET("/greet", controller.Greet)
	router.GET("/customers", cc.GetAllCustomers)
	router.GET("/customers/:customerID", cc.GetCustomer)
	router.POST("/customers/:customerID/account", ac.NewAccount)
	router.POST("/customers/:customerID/account/:accountID", ac.MakeTransaction)

}
func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:admin@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
