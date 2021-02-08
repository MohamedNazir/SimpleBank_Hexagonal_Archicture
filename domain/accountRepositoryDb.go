package domain

import (
	"fmt"
	"strconv"

	errs "github.com/MohamedNazir/SimpleBank/errors"
	"github.com/MohamedNazir/SimpleBank/logger"
	"github.com/jmoiron/sqlx"
)

const (
	QueryAccountInsert   = "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	QueryFindAccountByID = "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	QueryForWithdrawal   = "UPDATE accounts SET amount = amount - ? where account_id = ?"
	QueryForDeposit      = "UPDATE accounts SET amount = amount + ? where account_id = ?"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	result, err := d.client.Exec(QueryAccountInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error creating New Account" + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected Error from Database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retriving last insert ID for new Account" + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected Error from Database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) FindAccountByID(accountID string) (*Account, *errs.AppError) {
	var account Account
	err := d.client.Get(&account, QueryFindAccountByID, accountID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error Finding Account for %s %s", account.AccountID, err.Error()))
		return nil, errs.NewUnexpectedSystemError("Unexpected Error from Database")
	}
	return &account, nil
}

/**
 * transaction = make an entry in the transaction table + update the balance in the accounts table
 */
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {

	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.isWithdrawal() {
		_, err = tx.Exec(QueryForWithdrawal, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(QueryForDeposit, t.Amount, t.AccountID)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindAccountByID(t.AccountID)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionID = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}
