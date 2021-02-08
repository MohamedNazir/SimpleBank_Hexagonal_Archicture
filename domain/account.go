package domain

import (
	"github.com/MohamedNazir/SimpleBank/dto"
	errs "github.com/MohamedNazir/SimpleBank/errors"
)

//Account exported
type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

//AccountRepository interface
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindAccountByID(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

//ToNewAcountResponseDto exported
func (a Account) ToNewAcountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountID}
}

//CanWithdraw exported
func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
