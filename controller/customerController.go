package controller

import (
	"net/http"

	"github.com/MohamedNazir/SimpleBank/service"
	"github.com/gin-gonic/gin"
)

//CustomerController exported
type CustomerController struct {
	Service service.CustomerService
}

//GetAllCustomers exported
func (cc *CustomerController) GetAllCustomers(c *gin.Context) {

	customers, _ := cc.Service.GetAllCustomer()

	c.JSON(http.StatusOK, customers)
}
