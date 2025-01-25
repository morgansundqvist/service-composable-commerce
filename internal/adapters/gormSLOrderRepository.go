package adapters

import (
	"time"

	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
	"gorm.io/gorm"
)

type DBOrder struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	SessionId       string
	Email           string
	Name            string
	Address         string
	ZipCode         string
	City            string
	CompanyName     string
	Status          string
	CreatedDateTime time.Time
}

type DBOrderLine struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	OrderID   uuid.UUID
	ProductID uuid.UUID
	Price     int
	Quantity  int
}

type DBOrderLineContentLine struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	OrderLineID uuid.UUID
	ProductID   uuid.UUID
	Quantity    int
}

type GormSLOrderRepository struct {
	db *gorm.DB
}

func NewGormSLOrderRepository(db *gorm.DB) *GormSLOrderRepository {
	db.AutoMigrate(&DBOrder{}, &DBOrderLine{}, &DBOrderLineContentLine{})
	return &GormSLOrderRepository{db: db}
}

func toDBOrder(order *domain.Order) *DBOrder {
	return &DBOrder{
		ID:              order.ID,
		SessionId:       order.SessionId,
		Email:           order.Email,
		Name:            order.Name,
		Address:         order.Address,
		ZipCode:         order.ZipCode,
		City:            order.City,
		CompanyName:     order.CompanyName,
		Status:          order.Status,
		CreatedDateTime: order.CreatedDateTime,
	}
}

func toDomainOrder(dbOrder *DBOrder) *domain.Order {
	return &domain.Order{
		ID:              dbOrder.ID,
		SessionId:       dbOrder.SessionId,
		Email:           dbOrder.Email,
		Name:            dbOrder.Name,
		Address:         dbOrder.Address,
		ZipCode:         dbOrder.ZipCode,
		City:            dbOrder.City,
		CompanyName:     dbOrder.CompanyName,
		Status:          dbOrder.Status,
		CreatedDateTime: dbOrder.CreatedDateTime,
	}
}

func (r *GormSLOrderRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	dbOrder := toDBOrder(order)
	if err := r.db.Create(dbOrder).Error; err != nil {
		return nil, err
	}
	return toDomainOrder(dbOrder), nil
}

func (r *GormSLOrderRepository) UpdateOrder(order *domain.Order) error {
	dbOrder := toDBOrder(order)
	return r.db.Save(dbOrder).Error
}

func (r *GormSLOrderRepository) GetOrderById(id uuid.UUID) (*domain.Order, error) {
	var dbOrder DBOrder
	if err := r.db.Where("id = ?", id).First(&dbOrder).Error; err != nil {
		return nil, err
	}
	return toDomainOrder(&dbOrder), nil
}

func (r *GormSLOrderRepository) DeleteOrder(id uuid.UUID) error {
	return r.db.Delete(&DBOrder{}, id).Error
}

func (r *GormSLOrderRepository) CreateOrderLine(orderLine *domain.OrderLine) (*domain.OrderLine, error) {
	dbOrderLine := &DBOrderLine{
		ID:        orderLine.ID,
		OrderID:   orderLine.OrderID,
		ProductID: orderLine.ProductID,
		Price:     orderLine.Price,
		Quantity:  orderLine.Quantity,
	}
	if err := r.db.Create(dbOrderLine).Error; err != nil {
		return nil, err
	}
	return &domain.OrderLine{
		ID:        dbOrderLine.ID,
		OrderID:   dbOrderLine.OrderID,
		ProductID: dbOrderLine.ProductID,
		Price:     dbOrderLine.Price,
		Quantity:  dbOrderLine.Quantity,
	}, nil
}

func (r *GormSLOrderRepository) UpdateOrderLine(orderLine *domain.OrderLine) error {
	dbOrderLine := &DBOrderLine{
		ID:        orderLine.ID,
		OrderID:   orderLine.OrderID,
		ProductID: orderLine.ProductID,
		Price:     orderLine.Price,
		Quantity:  orderLine.Quantity,
	}
	return r.db.Save(dbOrderLine).Error
}

func (r *GormSLOrderRepository) GetOrderLineById(id uuid.UUID) (*domain.OrderLine, error) {
	var dbOrderLine DBOrderLine
	if err := r.db.Where("id = ?", id).First(&dbOrderLine).Error; err != nil {
		return nil, err
	}
	return &domain.OrderLine{
		ID:        dbOrderLine.ID,
		OrderID:   dbOrderLine.OrderID,
		ProductID: dbOrderLine.ProductID,
		Price:     dbOrderLine.Price,
		Quantity:  dbOrderLine.Quantity,
	}, nil
}

func (r *GormSLOrderRepository) DeleteOrderLine(id uuid.UUID) error {
	return r.db.Delete(&DBOrderLine{}, id).Error
}

func (r *GormSLOrderRepository) CreateOrderLineContentLine(contentLine *domain.OrderLineContentLine) (*domain.OrderLineContentLine, error) {
	dbOrderLineContentLine := &DBOrderLineContentLine{
		ID:          contentLine.ID,
		OrderLineID: contentLine.OrderLineID,
		ProductID:   contentLine.ProductID,
		Quantity:    contentLine.Quantity,
	}
	if err := r.db.Create(dbOrderLineContentLine).Error; err != nil {
		return nil, err
	}
	return &domain.OrderLineContentLine{
		ID:          dbOrderLineContentLine.ID,
		OrderLineID: dbOrderLineContentLine.OrderLineID,
		ProductID:   dbOrderLineContentLine.ProductID,
		Quantity:    dbOrderLineContentLine.Quantity,
	}, nil
}

func (r *GormSLOrderRepository) DeleteOrderLineContentLine(id uuid.UUID) error {
	return r.db.Delete(&DBOrderLineContentLine{}, id).Error
}
