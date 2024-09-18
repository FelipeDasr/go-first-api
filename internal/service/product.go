package service

import (
	"errors"
	"go-databases/internal/db"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

type CreateProductData struct {
	Name  string `json:"name" binding:"required"`
	Price int32 `json:"price" binding:"required,numeric,gt=0"`
	Stock int32 `json:"stock" binding:"required,numeric,gt=0"`
}

func (p *ProductService) CreateProduct(data *CreateProductData) (db.Product, error) {
	transaction, query, ctx, err := db.CreateQueryAndContextWithTx()

	if err != nil {
		return db.Product{}, errors.New("error when trying to create product")
	}
	defer transaction.Rollback()

	product, err := query.CreateProduct(ctx, db.CreateProductParams{
		Name:  data.Name,
		Price: data.Price,
		Stock: data.Stock,
	})

	if err != nil {
		return db.Product{}, errors.New("error creating product")
	}

	if err := transaction.Commit(); err != nil {
		return db.Product{}, errors.New("error whe trying to create product")
	}

	return product, nil;
}

func (p *ProductService) GetProductById(id int32) (db.Product, error) {
	query, ctx := db.CreateQueryAndContext()

	product, err := query.GetProductById(ctx, id)

	if err != nil {
		return db.Product{}, errors.New("product not found")
	}

	return product, nil;
}