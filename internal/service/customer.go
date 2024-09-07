package service

import (
	"errors"
	"go-databases/internal/db"
)

type CustomerServices struct {}

// NewCustomerServices creates a new CustomerServices
func NewCustomerServices() *CustomerServices {
	return &CustomerServices{}
}

type CreateCustomerData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// CreateCustomer creates a new customer
func (cs *CustomerServices) CreateCustomer(data *CreateCustomerData) (db.Customer, error) {
	query, ctx := db.CreateQueryAndContext()

	if _, err := query.CustomerAlreadyExistsByEmail(ctx, data.Email); err == nil {
		return db.Customer{}, errors.New("customer already exists with email: " + data.Email)
	}

	newCustomer, err := query.CreateCustomer(ctx, db.CreateCustomerParams{
		Name: data.Name,
		Email: data.Email,
	});

	if err != nil {
		return db.Customer{}, errors.New("error creating customer")
	}

	return newCustomer, nil
}

// GetCustomerById gets a customer by id
func (cs *CustomerServices) GetCustomerById(id int32) (db.Customer, error) {
	query, ctx := db.CreateQueryAndContext()

	customer, err := query.GetCustomerById(ctx, id)

	if err != nil {
		return db.Customer{}, errors.New("customer not found")
	}

	return customer, nil;
}