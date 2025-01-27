package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/morgansundqvist/service-composable-commerce/internal/application"
	"github.com/morgansundqvist/service-composable-commerce/internal/ports"
)

func SetupRouter(
	productService *application.ProductService,
	orderService *application.OrderService,
	engine *html.Engine,
	logger ports.Logger) *fiber.App {

	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/base",
		},
	)

	// Middleware to log URL and body if present
	// app.Use(func(c *fiber.Ctx) error {
	// 	if c.Body() != nil {
	// 		logger.Debug("Request", map[string]interface{}{
	// 			"URL":  c.OriginalURL(),
	// 			"Body": c.Body(),
	// 		})
	// 	}
	// 	return c.Next()
	// })

	productHandler := NewProductHandler(productService)
	orderHandler := NewOrderHandler(orderService)
	viewHandler := NewViewHandler(productService, orderService, logger)

	app.Get("/", viewHandler.HomePage)

	api := app.Group("/api")

	api.Post("/product-groups", productHandler.CreateProductGroup)
	api.Get("/product-groups", productHandler.GetProductGroups)
	api.Get("/product-groups/:id", productHandler.GetProductGroupByID)
	api.Get("/product-groups/:id/products", productHandler.GetProductsByProductGroupID)

	api.Post("/products", productHandler.CreateProduct)
	api.Get("/products", productHandler.GetProducts)
	api.Get("/products/:id", productHandler.GetProductByID)
	api.Patch("/products/:id", productHandler.UpdateProduct)
	api.Delete("/products/:id", productHandler.DeleteProduct)

	api.Post("/orders", orderHandler.CreateOrder)
	//api.Get("/orders", orderHandler.GetOrders) Maybe add later if needed
	api.Get("/orders/:id", orderHandler.GetOrderByID)
	api.Patch("/orders/:id", orderHandler.UpdateOrder)
	api.Delete("/orders/:id", orderHandler.DeleteOrder)

	api.Post("/sessions/:id", orderHandler.CreateSessionOrder)
	api.Get("/sessions/:id/order", orderHandler.GetOrderDetailsBySessionId)

	return app
}
