package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID `json:"id"`
	SessionId       string    `json:"session_id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	ZipCode         string    `json:"zip_code"`
	City            string    `json:"city"`
	CompanyName     string    `json:"company_name"`
	Status          string    `json:"status"`
	CreatedDateTime time.Time `json:"created_date_time"`
}

type CreateOrderInput struct {
	SessionId string `json:"session_id"`
}

type UpdateOrderInput struct {
	SessionId   *string `json:"session_id"`
	Email       *string `json:"email"`
	Name        *string `json:"name"`
	Address     *string `json:"address"`
	ZipCode     *string `json:"zip_code"`
	City        *string `json:"city"`
	CompanyName *string `json:"company_name"`
	Status      *string `json:"status"`
}

func CreateOrder(input CreateOrderInput) (*Order, error) {
	order := &Order{
		SessionId:       input.SessionId,
		CreatedDateTime: time.Now(),
	}

	return order, nil
}

func (o *Order) Update(input UpdateOrderInput) error {
	if input.SessionId != nil {
		o.SessionId = *input.SessionId
	}
	if input.Email != nil {
		o.Email = *input.Email
	}
	if input.Name != nil {
		o.Name = *input.Name
	}
	if input.Address != nil {
		o.Address = *input.Address
	}
	if input.ZipCode != nil {
		o.ZipCode = *input.ZipCode
	}
	if input.City != nil {
		o.City = *input.City
	}
	if input.CompanyName != nil {
		o.CompanyName = *input.CompanyName
	}
	if input.Status != nil {
		o.Status = *input.Status
	}

	return nil
}

func (o *Order) SetStatus(status string) error {
	o.Status = status

	return nil
}

type OrderLine struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
}

type OrderLineContentLine struct {
	ID          uuid.UUID `json:"id"`
	OrderLineID uuid.UUID `json:"order_line_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
}

type CreateOrderLineInput struct {
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
}

type UpdateOrderLineInput struct {
	OrderID   *uuid.UUID `json:"order_id"`
	ProductID *uuid.UUID `json:"product_id"`
	Price     *int       `json:"price"`
	Quantity  *int       `json:"quantity"`
}

type CreateOrderLineContentLineInput struct {
	OrderLineID uuid.UUID `json:"order_line_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
}

func CreateOrderLine(input CreateOrderLineInput) (*OrderLine, error) {
	orderLine := &OrderLine{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Price:     input.Price,
		Quantity:  input.Quantity,
	}

	return orderLine, nil
}

func (ol *OrderLine) UpdateQuantity(quantity int) error {
	ol.Quantity = quantity

	return nil
}

func CreateOrderLineContentLine(input CreateOrderLineContentLineInput) (*OrderLineContentLine, error) {
	orderLineContentLine := &OrderLineContentLine{
		OrderLineID: input.OrderLineID,
		ProductID:   input.ProductID,
		Quantity:    input.Quantity,
	}

	return orderLineContentLine, nil
}
