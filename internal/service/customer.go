package service

import (
	"context"
	"database/sql"
	"errors"
	"go-databases/internal/db"
)

type CustomerServices struct {
	Db *sql.DB
}

// NewCustomerServices creates a new CustomerServices
func NewCustomerServices(db *sql.DB) *CustomerServices {
	return &CustomerServices{Db: db}
}

type CreateCustomerData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (cs *CustomerServices) CreateCustomer(data *CreateCustomerData) (db.Customer, error) {
	ctx := context.Background()
	query := db.New(cs.Db)

	if _, err := query.CustomerAlreadyExistsByEmail(ctx, data.Email); err == nil {
		return db.Customer{}, errors.New("Customer already exists with email: " + data.Email)
	}

	return query.CreateCustomer(ctx, db.CreateCustomerParams{
		Name: data.Name,
		Email: data.Email,
	});
}