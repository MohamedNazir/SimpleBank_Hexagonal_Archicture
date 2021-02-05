package domain

import (
	"database/sql"
	"time"

	errs "github.com/MohamedNazir/SimpleBank/errors"
	"github.com/MohamedNazir/SimpleBank/logger"
	_ "github.com/go-sql-driver/mysql"
)

//CustomerRepositoryDb exported
type CustomerRepositoryDb struct {
	client *sql.DB
}

const (
	queryFindAll  = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	queryFindByID = "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"
)

//FindAll exported
func (db CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	rows, err := db.client.Query(queryFindAll)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedSystemError("Unexpected System error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedSystemError("Unexpected System error")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

//FindByID exported
func (db CustomerRepositoryDb) FindByID(id int) (*Customer, *errs.AppError) {

	row := db.client.QueryRow(queryFindByID, id)

	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Record Not Found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedSystemError("Unexpected System Error")
		}

	}
	return &c, nil

}

//NewCustomerRepositoryDb exported
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
