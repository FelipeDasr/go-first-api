package service

import (
	"errors"
	"go-databases/internal/db"
)

type OrderService struct{}

type CreateOrderData struct {
	CustomerID  int32 `json:"customer_id" binding:"required,numeric"`
	ProductID   int32 `json:"product_id" binding:"required,numeric"`
	UnitsAmount int32 `json:"units_amount" binding:"required,numeric,gt=0"`
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) CreateOrder(data *CreateOrderData) (db.Order, error) {
	transaction, query, ctx, err := db.CreateQueryAndContextWithTx()

	if err != nil {
		return db.Order{}, errors.New("error when trying to create order")
	}
	defer transaction.Rollback()

	product, err := query.GetProductById(ctx, data.ProductID);

	if err != nil {
		return db.Order{}, errors.New("product " + string(data.ProductID) + " not found")
	}

	if product.Stock < data.UnitsAmount {
		return db.Order{}, errors.New("not enough stock")
	}

	if err := query.IncrementProductStockById(ctx, db.IncrementProductStockByIdParams{
		ID: data.ProductID,
		Stock: -1 * data.UnitsAmount,
	});err != nil {
		return db.Order{}, errors.New("error when trying to create order")
	}

	order, err := query.CreateOrder(ctx, db.CreateOrderParams{
		CustomerID: data.CustomerID,
		ProductID: data.ProductID,
		UnitsAmount: data.UnitsAmount,
		UnitPrice: product.Price,
	})

	if err != nil {
		return db.Order{}, errors.New("error creating order")
	}

	if err := transaction.Commit(); err != nil {
		return db.Order{}, errors.New("error whe trying to create order")
	}

	return order, nil
}