package ports

import (
	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
)

type ProductRepository interface {
	// CreateProduct creates a new product
	CreateProduct(product *domain.Product) error
	// UpdateProduct updates a product
	UpdateProduct(product *domain.Product) error
	// DeleteProduct deletes a product
	DeleteProduct(productID uuid.UUID) error
	// GetProduct retrieves a product by its ID
	GetProduct(productID uuid.UUID) (*domain.Product, error)
	// ListProducts retrieves all products
	ListProducts() ([]domain.Product, error)
	// CreateProductGroup creates a new product group
	CreateProductGroup(productGroup *domain.ProductGroup) error
	// UpdateProductGroup updates a product group
	UpdateProductGroup(productGroup *domain.ProductGroup) error
	// DeleteProductGroup deletes a product group
	DeleteProductGroup(productGroupID uuid.UUID) error
	// GetProductGroup retrieves a product group by its ID
	GetProductGroup(productGroupID uuid.UUID) (*domain.ProductGroup, error)
	// ListProductGroups retrieves all product groups
	ListProductGroups() ([]domain.ProductGroup, error)
}
