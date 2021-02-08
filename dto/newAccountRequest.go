package dto

import (
	"strings"

	errs "github.com/MohamedNazir/SimpleBank/errors"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customerId"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Minimum amount to deposit is 5000.00")
	}
	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be savings or checking")
	}
	return nil
}

type NewAccountResponse struct {
	AccountID string `json:"accountId`
}
