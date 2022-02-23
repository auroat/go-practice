package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Customer Class
type Customer struct {
	CustomerId
	int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver,
		databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

//GetCustomerById with parameter customerId returns Customer
func GetCustomerById(customerId int) Customer {
	var database *sql.DB
	database = GetConnection()
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}

	return customer
}
