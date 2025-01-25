package application

import (
	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
	"github.com/morgansundqvist/service-composable-commerce/internal/ports"
)

type ProductService struct {
	productRepository ports.ProductRepository
	logger            ports.Logger
}

func NewProductService(productRepository ports.ProductRepository, logger ports.Logger) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		logger:            logger,
	}
}

func (s *ProductService) GetProductGroups() ([]domain.ProductGroup, error) {
	productGroups, err := s.productRepository.ListProductGroups()
	if err != nil {
		s.logger.Error("failed to get product groups", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return productGroups, nil
}

func (s *ProductService) CreateProductGroup(productGroupInput domain.CreateProductGroupInput) (*domain.ProductGroup, error) {
	productGroup, err := domain.CreateProductGroup(productGroupInput)
	if err != nil {
		return nil, err
	}

	err = s.productRepository.CreateProductGroup(productGroup)

	if err != nil {
		s.logger.Error("failed to create product group", map[string]interface{}{
			"error": err,
		})

		return nil, err
	}

	return productGroup, nil
}

func (s *ProductService) GetProductGroupByID(id string) (*domain.ProductGroup, error) {

	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})

		return nil, err
	}

	productGroup, err := s.productRepository.GetProductGroup(uuidId)
	if err != nil {
		s.logger.Error("failed to get product group by ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return productGroup, nil
}

func (s *ProductService) CreateProduct(productInput domain.CreateProductInput) (*domain.Product, error) {
	product, err := domain.CreateProduct(productInput)
	if err != nil {
		return nil, err
	}

	err = s.productRepository.CreateProduct(product)
	if err != nil {
		s.logger.Error("failed to create product", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id string, productInput domain.UpdateProductInput) (*domain.Product, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	product, err := s.productRepository.GetProduct(uuidId)
	if err != nil {
		s.logger.Error("failed to get product by ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	err = product.Update(productInput)
	if err != nil {
		return nil, err
	}

	err = s.productRepository.UpdateProduct(product)
	if err != nil {
		s.logger.Error("failed to update product", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return err
	}

	err = s.productRepository.DeleteProduct(uuidId)
	if err != nil {
		s.logger.Error("failed to delete product", map[string]interface{}{
			"error": err,
		})
		return err
	}

	return nil
}

func (s *ProductService) GetProductByID(id string) (*domain.Product, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("failed to parse UUID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	product, err := s.productRepository.GetProduct(uuidId)
	if err != nil {
		s.logger.Error("failed to get product by ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return product, nil
}

func (s *ProductService) ListProducts() ([]domain.Product, error) {
	products, err := s.productRepository.ListProducts()
	if err != nil {
		s.logger.Error("failed to list products", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return products, nil
}
