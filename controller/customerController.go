package controller

import (
	"net/http"
	"strconv"

	"github.com/MohamedNazir/SimpleBank/service"
	"github.com/gin-gonic/gin"
)

//CustomerController exported
type CustomerController struct {
	Service service.CustomerService
}

//GetAllCustomers exported
func (cc *CustomerController) GetAllCustomers(c *gin.Context) {

	customers, err := cc.Service.GetAllCustomer()
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}

	c.JSON(http.StatusOK, customers)
}

//GetAllCustomers exported
func (cc *CustomerController) GetCustomer(c *gin.Context) {

	customerID, usrErr := getParameter(c.Param("customerID"))
	if usrErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid customer ID")
		return
	}
	customers, err := cc.Service.GetCustomer(customerID)
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getParameter(userIDParam string) (int, error) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {

		return 0, err
	}
	return int(userID), nil
}
