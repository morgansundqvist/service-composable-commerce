package adapters

import (
	"github.com/google/uuid"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
	"gorm.io/gorm"
)

type DBProduct struct {
	ID                         uuid.UUID `gorm:"type:uuid;primary_key"`
	Name                       string
	Price                      int
	ProductGroupID             uuid.UUID
	Order                      int
	IsConfigurable             bool
	ConfiguredByProductGroupID *uuid.UUID
	ConfiguredQuantity         int
	IsSoldSeparately           bool
}

type DBProductGroup struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Name   string
	Order  int
	IsSold bool
}

type GormSLProductRepository struct {
	db *gorm.DB
}

func NewGormSLProductRepository(db *gorm.DB) *GormSLProductRepository {
	db.AutoMigrate(&DBProduct{}, &DBProductGroup{})
	return &GormSLProductRepository{db: db}
}

func toDBProduct(product *domain.Product) *DBProduct {
	return &DBProduct{
		ID:                         product.ID,
		Name:                       product.Name,
		Price:                      product.Price,
		ProductGroupID:             product.ProductGroupID,
		Order:                      product.Order,
		IsConfigurable:             product.IsConfigurable,
		ConfiguredByProductGroupID: product.ConfiguredByProductGroupID,
		ConfiguredQuantity:         product.ConfiguredQuantity,
		IsSoldSeparately:           product.IsSoldSeparately,
	}
}

func toDomainProduct(dbProduct *DBProduct) *domain.Product {
	return &domain.Product{
		ID:                         dbProduct.ID,
		Name:                       dbProduct.Name,
		Price:                      dbProduct.Price,
		ProductGroupID:             dbProduct.ProductGroupID,
		Order:                      dbProduct.Order,
		IsConfigurable:             dbProduct.IsConfigurable,
		ConfiguredByProductGroupID: dbProduct.ConfiguredByProductGroupID,
		ConfiguredQuantity:         dbProduct.ConfiguredQuantity,
		IsSoldSeparately:           dbProduct.IsSoldSeparately,
	}
}

func toDBProductGroup(productGroup *domain.ProductGroup) *DBProductGroup {
	return &DBProductGroup{
		ID:     productGroup.ID,
		Name:   productGroup.Name,
		Order:  productGroup.Order,
		IsSold: productGroup.IsSold,
	}
}

func toDomainProductGroup(dbProductGroup *DBProductGroup) *domain.ProductGroup {
	return &domain.ProductGroup{
		ID:     dbProductGroup.ID,
		Name:   dbProductGroup.Name,
		Order:  dbProductGroup.Order,
		IsSold: dbProductGroup.IsSold,
	}
}

func (r *GormSLProductRepository) CreateProduct(product *domain.Product) error {
	dbProduct := toDBProduct(product)
	return r.db.Create(dbProduct).Error
}

func (r *GormSLProductRepository) UpdateProduct(product *domain.Product) error {
	dbProduct := toDBProduct(product)
	return r.db.Save(dbProduct).Error
}

func (r *GormSLProductRepository) DeleteProduct(productID uuid.UUID) error {
	return r.db.Delete(&DBProduct{}, productID).Error
}

func (r *GormSLProductRepository) GetProduct(productID uuid.UUID) (*domain.Product, error) {
	var dbProduct DBProduct
	err := r.db.Where("id = ?", productID).First(&dbProduct).Error
	if err != nil {
		return nil, err
	}
	return toDomainProduct(&dbProduct), nil
}

func (r *GormSLProductRepository) ListProducts() ([]domain.Product, error) {
	var dbProducts []DBProduct
	err := r.db.Find(&dbProducts).Error
	if err != nil {
		return nil, err
	}
	products := make([]domain.Product, len(dbProducts))
	for i, dbProduct := range dbProducts {
		products[i] = *toDomainProduct(&dbProduct)
	}
	return products, nil
}

func (r *GormSLProductRepository) CreateProductGroup(productGroup *domain.ProductGroup) error {
	dbProductGroup := toDBProductGroup(productGroup)
	return r.db.Create(dbProductGroup).Error
}

func (r *GormSLProductRepository) UpdateProductGroup(productGroup *domain.ProductGroup) error {
	dbProductGroup := toDBProductGroup(productGroup)
	return r.db.Save(dbProductGroup).Error
}

func (r *GormSLProductRepository) DeleteProductGroup(productGroupID uuid.UUID) error {
	return r.db.Delete(&DBProductGroup{}, productGroupID).Error
}

func (r *GormSLProductRepository) GetProductGroup(productGroupID uuid.UUID) (*domain.ProductGroup, error) {
	var dbProductGroup DBProductGroup
	err := r.db.Where("id = ?", productGroupID).First(&dbProductGroup).Error
	if err != nil {
		return nil, err
	}
	return toDomainProductGroup(&dbProductGroup), nil
}

func (r *GormSLProductRepository) ListProductGroups() ([]domain.ProductGroup, error) {
	var dbProductGroups []DBProductGroup
	err := r.db.Find(&dbProductGroups).Error
	if err != nil {
		return nil, err
	}
	productGroups := make([]domain.ProductGroup, len(dbProductGroups))
	for i, dbProductGroup := range dbProductGroups {
		productGroups[i] = *toDomainProductGroup(&dbProductGroup)
	}
	return productGroups, nil
}
