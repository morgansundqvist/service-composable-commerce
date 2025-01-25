package main

import (
	"github.com/morgansundqvist/service-composable-commerce/internal/adapters"
	"github.com/morgansundqvist/service-composable-commerce/internal/api"
	"github.com/morgansundqvist/service-composable-commerce/internal/application"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logger := adapters.NewLogrusLogger()

	db, err := gorm.Open(sqlite.Open("commerce.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database", map[string]interface{}{
			"error": err,
		})
	}

	productRepository := adapters.NewGormSLProductRepository(db)
	orderRepository := adapters.NewGormSLOrderRepository(db)

	productService := application.NewProductService(productRepository, logger)
	productHandler := api.NewProductHandler(productService)

	orderService := application.NewOrderService(orderRepository, logger)
	orderHandler := api.NewOrderHandler(orderService)

	app := api.SetupRouter(productHandler, orderHandler)

	app.Listen(":3000")
}
