package handler

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"sample/model"
	"sample/store"
)

type handler struct {
	store store.Product
}

func New(s store.Product) handler {
	return handler{store: s}
}

// Create to create new product
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var product model.Product

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("Error in binding: %v", err)

		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, &product)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetByID to get product by ID
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// ctx.PathParam() returns the path parameter from HTTP request.
	id, err := validateID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	resp, err := h.store.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update to update product by ID
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	id, err := validateID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	var product model.Product
	if err = ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("Error in binding: %v", err)

		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	product.ID = id

	resp, err := h.store.Update(ctx, &product)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Delete to delete product by ID
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id, err := validateID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	return nil, h.store.Delete(ctx, id)
}

func validateID(id string) (int, error) {
	if id == "" {
		return 0, errors.MissingParam{Param: []string{"id"}}
	}

	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.InvalidParam{Param: []string{"id"}}
	}

	return res, err
}

// GetAll to get all products
func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	return h.store.GetAll(ctx)
}
