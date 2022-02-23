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

	defer database.Close()
	return customer
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}
	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}
