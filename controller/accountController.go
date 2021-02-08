package controller

import (
	"net/http"
	"strconv"

	"github.com/MohamedNazir/SimpleBank/dto"
	"github.com/MohamedNazir/SimpleBank/service"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service service.AccountService
}

func (ac AccountController) NewAccount(c *gin.Context) {

	var acc dto.NewAccountRequest
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	customerID, usrErr := getParameter(c.Param("customerID"))
	if usrErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid customer ID")
		return
	}

	acc.CustomerID = strconv.Itoa(customerID)
	resp, err := ac.Service.NewAccount(acc)
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}
	c.JSON(http.StatusCreated, resp)

}

func (ac AccountController) MakeTransaction(c *gin.Context) {

	var txReq dto.TransactionRequest
	if err := c.ShouldBindJSON(&txReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	customerID, usrErr := getParameter(c.Param("customerID"))
	if usrErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid customer ID")
		return
	}

	accountID, usrErr := getParameter(c.Param("accountID"))
	if usrErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid account ID")
		return
	}

	txReq.CustomerID = strconv.Itoa(customerID)
	txReq.AccountID = strconv.Itoa(accountID)

	resp, err := ac.Service.MakeTransaction(txReq)
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}
	c.JSON(http.StatusCreated, resp)

}
