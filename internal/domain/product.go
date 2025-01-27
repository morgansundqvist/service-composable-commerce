package domain

import (
	"errors"

	"github.com/google/uuid"
)

type ProductGroup struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Order  int       `json:"order"`
	IsSold bool      `json:"is_sold"`
}

type CreateProductGroupInput struct {
	Name   string `json:"name"`
	Order  int    `json:"order"`
	IsSold bool   `json:"is_sold"`
}

type UpdateProductGroupInput struct {
	Name   *string `json:"name"`
	Order  *int    `json:"order"`
	IsSold *bool   `json:"is_sold"`
}

type Product struct {
	ID                         uuid.UUID  `json:"id"`
	Name                       string     `json:"name"`
	Price                      int        `json:"price"`
	ProductGroupID             uuid.UUID  `json:"product_group_id"`
	Order                      int        `json:"order"`
	IsConfigurable             bool       `json:"is_configurable"`
	ConfiguredByProductGroupID *uuid.UUID `json:"configured_by_product_group_id"`
	ConfiguredQuantity         int        `json:"configured_quantity"`
	IsSoldSeparately           bool       `json:"is_sold_separately"`
}

type CreateProductInput struct {
	Name                       string     `json:"name"`
	Price                      int        `json:"price"`
	ProductGroupID             uuid.UUID  `json:"product_group_id"`
	Order                      int        `json:"order"`
	IsConfigurable             bool       `json:"is_configurable"`
	ConfiguredByProductGroupID *uuid.UUID `json:"configured_by_product_group_id"`
	ConfiguredQuantity         int        `json:"configured_quantity"`
	IsSoldSeparately           bool       `json:"is_sold_separately"`
}

// UpdateProductInput defines the data required to update an existing product
type UpdateProductInput struct {
	Name                       *string    `json:"name"`
	Price                      *int       `json:"price"`
	Order                      *int       `json:"order"`
	IsConfigurable             *bool      `json:"is_configurable"`
	ConfiguredByProductGroupID *uuid.UUID `json:"configured_by_product_group_id"`
	ConfiguredQuantity         *int       `json:"configured_quantity"`
	IsSoldSeparately           *bool      `json:"is_sold_separately"`
}

func CreateProduct(input CreateProductInput) (*Product, error) {
	if input.Name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if input.Price < 0 {
		return nil, errors.New("product price cannot be negative")
	}
	if input.ConfiguredQuantity < 0 {
		return nil, errors.New("configured quantity cannot be negative")
	}

	product := Product{
		ID:                         uuid.New(),
		Name:                       input.Name,
		Price:                      input.Price,
		ProductGroupID:             input.ProductGroupID,
		Order:                      input.Order,
		IsConfigurable:             input.IsConfigurable,
		ConfiguredByProductGroupID: input.ConfiguredByProductGroupID,
		ConfiguredQuantity:         input.ConfiguredQuantity,
		IsSoldSeparately:           input.IsSoldSeparately,
	}

	return &product, nil
}

func (p *Product) Update(input UpdateProductInput) error {
	if input.Name != nil {
		p.Name = *input.Name
	}
	if input.Price != nil {
		p.Price = *input.Price
	}
	if input.Order != nil {
		p.Order = *input.Order
	}
	if input.IsConfigurable != nil {
		p.IsConfigurable = *input.IsConfigurable
	}
	if input.ConfiguredByProductGroupID != nil {
		p.ConfiguredByProductGroupID = input.ConfiguredByProductGroupID
	}
	if input.ConfiguredQuantity != nil {
		p.ConfiguredQuantity = *input.ConfiguredQuantity
	}
	if input.IsSoldSeparately != nil {
		p.IsSoldSeparately = *input.IsSoldSeparately
	}

	return nil
}

func CreateProductGroup(input CreateProductGroupInput) (*ProductGroup, error) {
	if input.Name == "" {
		return nil, errors.New("product group name cannot be empty")
	}

	productGroup := ProductGroup{
		ID:     uuid.New(),
		Name:   input.Name,
		Order:  input.Order,
		IsSold: input.IsSold,
	}

	return &productGroup, nil
}

func (pg *ProductGroup) Update(input UpdateProductGroupInput) error {
	if input.Name != nil {
		pg.Name = *input.Name
	}
	if input.Order != nil {
		pg.Order = *input.Order
	}
	if input.IsSold != nil {
		pg.IsSold = *input.IsSold
	}

	return nil
}

type ProductGroupWithProducts struct {
	ProductGroup ProductGroup `json:"product_group"`
	Products     []Product    `json:"products"`
}
