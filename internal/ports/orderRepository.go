package ports

import (
	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(order *domain.Order) error
	GetOrderById(id uuid.UUID) (*domain.Order, error)
	DeleteOrder(id uuid.UUID) error
	CreateOrderLine(orderLine *domain.OrderLine) (*domain.OrderLine, error)
	UpdateOrderLine(orderLine *domain.OrderLine) error
	GetOrderLineById(id uuid.UUID) (*domain.OrderLine, error)
	DeleteOrderLine(id uuid.UUID) error
	CreateOrderLineContentLine(contentLine *domain.OrderLineContentLine) (*domain.OrderLineContentLine, error)
	DeleteOrderLineContentLine(id uuid.UUID) error
	GetOrderBySessionId(sessionId string) (*domain.Order, error)
	GetOrderLinesByOrderId(orderId uuid.UUID) ([]*domain.OrderLine, error)
	GetOrderLineContentLinesByOrderLineId(orderLineId uuid.UUID) ([]*domain.OrderLineContentLine, error)
	GetOrderByStatus(status string) ([]*domain.Order, error)
}
