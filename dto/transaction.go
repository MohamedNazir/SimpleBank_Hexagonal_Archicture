package dto

import (
	errs "github.com/MohamedNazir/SimpleBank/errors"
)

const (
	//WITHDRAWAL exported
	WITHDRAWAL = "withdrawal"
	//DEPOSIT exported
	DEPOSIT = "deposit"
)

//TransactionRequest dto request struct
type TransactionRequest struct {
	AccountID       string  `json:"accountId"`
	CustomerID      string  `json:"customerId"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transactionType"`
	TransactionDate string  `json:"transactionDate"`
	AccountType     string  `json:"accountType"`
}

// Validate used to validate the incoming request
func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

//IsTransactionTypeWithdrawal exported
func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

//IsTransactionTypeDeposit exported
func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

//TransactionResponse dto response struct
type TransactionResponse struct {
	TransactionID   string  `json:"transactionId"`
	AccountID       string  `json:"accountId"`
	Amount          float64 `json:"newBalance"`
	TransactionType string  `json:"transactionType"`
	TransactionDate string  `json:"transactionDate"`
}
