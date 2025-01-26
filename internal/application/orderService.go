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

func (s *OrderService) GetOrderDetailsBySessionId(sessionId string) (*DTOOrderDetails, error) {
	order, err := s.orderRepository.GetOrderBySessionId(sessionId)
	if err != nil {
		s.logger.Error("failed to get order by session ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	orderLines, err := s.orderRepository.GetOrderLinesByOrderId(order.ID)
	if err != nil {
		s.logger.Error("failed to get order lines by order ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	dtoOrderLines := make([]DTOOrderLine, len(orderLines))
	for i, orderLine := range orderLines {
		contentLines, err := s.orderRepository.GetOrderLineContentLinesByOrderLineId(orderLine.ID)
		if err != nil {
			s.logger.Error("failed to get order line content lines by order line ID", map[string]interface{}{
				"error": err,
			})
			return nil, err
		}

		dtoContentLines := make([]DTOOrderLineContentLine, len(contentLines))
		for j, contentLine := range contentLines {
			dtoContentLines[j] = DTOOrderLineContentLine{
				ID:          contentLine.ID,
				OrderLineID: contentLine.OrderLineID,
				ProductID:   contentLine.ProductID,
				Product:     domain.Product{}, // Assuming you have a way to get the product details
				Quantity:    contentLine.Quantity,
			}
		}

		dtoOrderLines[i] = DTOOrderLine{
			ID:           orderLine.ID,
			OrderID:      orderLine.OrderID,
			ProductID:    orderLine.ProductID,
			Product:      domain.Product{}, // Assuming you have a way to get the product details
			ContentLines: dtoContentLines,
		}
	}

	dtoOrder := DTOOrder{
		ID:              order.ID,
		SessionID:       order.SessionId,
		Email:           order.Email,
		Name:            order.Name,
		Address:         order.Address,
		ZipCode:         order.ZipCode,
		City:            order.City,
		CompanyName:     order.CompanyName,
		Status:          order.Status,
		CreatedDateTime: order.CreatedDateTime.String(),
		OrderLines:      dtoOrderLines,
	}

	return &DTOOrderDetails{Order: dtoOrder}, nil
}

type DTOOrderDetails struct {
	Order DTOOrder `json:"order"`
}

type DTOOrder struct {
	ID              uuid.UUID      `json:"id"`
	SessionID       string         `json:"session_id"`
	Email           string         `json:"email"`
	Name            string         `json:"name"`
	Address         string         `json:"address"`
	ZipCode         string         `json:"zip_code"`
	City            string         `json:"city"`
	CompanyName     string         `json:"company_name"`
	Status          string         `json:"status"`
	CreatedDateTime string         `json:"created_date_time"`
	OrderLines      []DTOOrderLine `json:"order_lines"`
}

type DTOOrderLine struct {
	ID           uuid.UUID                 `json:"id"`
	OrderID      uuid.UUID                 `json:"order_id"`
	ProductID    uuid.UUID                 `json:"product_id"`
	Product      domain.Product            `json:"product"`
	ContentLines []DTOOrderLineContentLine `json:"content_lines"`
}

type DTOOrderLineContentLine struct {
	ID          uuid.UUID      `json:"id"`
	OrderLineID uuid.UUID      `json:"order_line_id"`
	ProductID   uuid.UUID      `json:"product_id"`
	Product     domain.Product `json:"product"`
	Quantity    int            `json:"quantity"`
}
