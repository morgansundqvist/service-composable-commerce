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

type DTOProductGroupDetails struct {
	ProductGroup domain.ProductGroup `json:"product_group"`
}

type DTOProductDetails struct {
	Product domain.Product `json:"product"`
}

type DTOProductGroupList struct {
	ProductGroups []domain.ProductGroup `json:"product_groups"`
}

type DTOProductList struct {
	Products []domain.Product `json:"products"`
}

type DTOProductGroupWithProducts struct {
	ProductGroups []domain.ProductGroupWithProducts `json:"product_groups"`
}

func NewProductService(productRepository ports.ProductRepository, logger ports.Logger) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		logger:            logger,
	}
}

func (s *ProductService) GetProductGroups() (*DTOProductGroupList, error) {
	productGroups, err := s.productRepository.ListProductGroups()
	if err != nil {
		s.logger.Error("failed to get product groups", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return &DTOProductGroupList{ProductGroups: productGroups}, nil
}

func (s *ProductService) CreateProductGroup(productGroupInput domain.CreateProductGroupInput) (*DTOProductGroupDetails, error) {
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

	return &DTOProductGroupDetails{ProductGroup: *productGroup}, nil
}

func (s *ProductService) GetProductGroupByID(id string) (*DTOProductGroupDetails, error) {
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

	return &DTOProductGroupDetails{ProductGroup: *productGroup}, nil
}

func (s *ProductService) CreateProduct(productInput domain.CreateProductInput) (*DTOProductDetails, error) {
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

	return &DTOProductDetails{Product: *product}, nil
}

func (s *ProductService) UpdateProduct(id string, productInput domain.UpdateProductInput) (*DTOProductDetails, error) {
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

	return &DTOProductDetails{Product: *product}, nil
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

func (s *ProductService) GetProductByID(id string) (*DTOProductDetails, error) {
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

	return &DTOProductDetails{Product: *product}, nil
}

func (s *ProductService) ListProducts() (*DTOProductList, error) {
	products, err := s.productRepository.ListProducts()
	if err != nil {
		s.logger.Error("failed to list products", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return &DTOProductList{Products: products}, nil
}

func (s *ProductService) GetProductsByProductGroupID(id uuid.UUID) (*DTOProductList, error) {
	products, err := s.productRepository.ListProductsByProductGroupID(id)
	if err != nil {
		s.logger.Error("failed to list products by product group ID", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return &DTOProductList{Products: products}, nil
}

func (s *ProductService) GetProductGroupsWithProducts() (*DTOProductGroupWithProducts, error) {
	productGroupsWithProducts, err := s.productRepository.ListProductGroupsWithProducts()
	if err != nil {
		s.logger.Error("failed to get product groups with products", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return &DTOProductGroupWithProducts{ProductGroups: productGroupsWithProducts}, nil
}
