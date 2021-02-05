package api

import (
	"github.com/MohamedNazir/SimpleBank/controller"
	"github.com/MohamedNazir/SimpleBank/domain"
	"github.com/MohamedNazir/SimpleBank/service"
)

func mapUrls() {

	// wiring
	//cc := controller.CustomerController{Service: service.CustomerService(domain.NewCustomerRepositoryStub())}
	//	cc := controller.CustomerController{Service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := controller.CustomerController{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.GET("/greet", controller.Greet)
	router.GET("/customers", ch.GetAllCustomers)
	router.GET("/customers/:customerID", ch.GetCustomer)

}
