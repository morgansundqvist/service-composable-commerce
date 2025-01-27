package main

import (
	"time"

	"github.com/gofiber/template/html/v2"
	"github.com/morgansundqvist/service-composable-commerce/internal/adapters"
	"github.com/morgansundqvist/service-composable-commerce/internal/api"
	"github.com/morgansundqvist/service-composable-commerce/internal/application"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logger := adapters.NewLogrusLogger()
	logger.SetLogLevel("debug")

	db, err := gorm.Open(sqlite.Open("commerce.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database", map[string]interface{}{
			"error": err,
		})
	}

	productRepository := adapters.NewGormSLProductRepository(db, logger)
	orderRepository := adapters.NewGormSLOrderRepository(db)

	productService := application.NewProductService(productRepository, logger)
	orderService := application.NewOrderService(orderRepository, logger)

	// Setup the template engine
	engine := html.New("./views", ".html")

	app := api.SetupRouter(productService, orderService, engine, logger)

	//run delete order job every 5 minutes
	go func() {
		for {
			logger.Info("starting execution of delete old orders job", nil)
			orderService.RemoveOldCreatedOrders()
			<-time.After(5 * time.Minute)
		}
	}()

	app.Listen(":3000")
}
