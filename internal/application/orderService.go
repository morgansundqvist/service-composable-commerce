package application

import (
	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
	"github.com/morgansundqvist/service-composable-commerce/internal/ports"
)

type OrderService struct {
	orderRepository ports.OrderRepository
	logger          ports.Logger
}

func NewOrderService(orderRepository ports.OrderRepository, logger ports.Logger) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		logger:          logger,
	}
}

func (s *OrderService) CreateOrder(input domain.CreateOrderInput) (*domain.Order, error) {
	order, err := domain.CreateOrder(input)
	if err != nil {
		return nil, err
	}

	order, err = s.orderRepository.CreateOrder(order)
	if err != nil {
		s.logger.Error("failed to create order", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(id string, input domain.UpdateOrderInput) (*domain.Order, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	order, err := s.orderRepository.GetOrderById(uuidId)
	if err != nil {
		s.logger.Error("failed to get order by ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	err = order.Update(input)
	if err != nil {
		return nil, err
	}

	err = s.orderRepository.UpdateOrder(order)
	if err != nil {
		s.logger.Error("failed to update order", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrderById(id string) (*domain.Order, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	order, err := s.orderRepository.GetOrderById(uuidId)
	if err != nil {
		s.logger.Error("failed to get order by ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return order, nil
}

func (s *OrderService) DeleteOrder(id string) error {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return err
	}

	err = s.orderRepository.DeleteOrder(uuidId)
	if err != nil {
		s.logger.Error("failed to delete order", map[string]interface{}{
			"error": err,
		})
		return err
	}

	return nil
}
