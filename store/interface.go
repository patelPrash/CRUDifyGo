package store

import (
	"sample/model"

	"gofr.dev/pkg/gofr"
)

type Product interface {
	// Create inserts a new product record into the database
	Create(ctx *gofr.Context, product *model.Product) (*model.Product, error)

	// GetByID retrieves a Product record based on its ID
	GetByID(ctx *gofr.Context, id int) (*model.Product, error)

	// Update updates an existing product record with the provided information
	Update(ctx *gofr.Context, product *model.Product) (*model.Product, error)

	// Delete removes a product record from the database based on its ID
	Delete(ctx *gofr.Context, id int) error

	// GetAll retrieves all product records from the database
	GetAll(ctx *gofr.Context) ([]*model.Product, error)
}
