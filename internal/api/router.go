package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(productHandler *ProductHandler, orderHandler *OrderHandler) *fiber.App {
	app := fiber.New()

	app.Post("/api/product-groups", productHandler.CreateProductGroup)
	app.Get("/api/product-groups", productHandler.GetProductGroups)
	app.Get("/api/product-groups/:id", productHandler.GetProductGroupByID)

	app.Post("/api/products", productHandler.CreateProduct)
	app.Get("/api/products", productHandler.GetProducts)
	app.Get("/api/products/:id", productHandler.GetProductByID)
	app.Patch("/api/products/:id", productHandler.UpdateProduct)
	app.Delete("/api/products/:id", productHandler.DeleteProduct)

	app.Post("/api/orders", orderHandler.CreateOrder)
	//app.Get("/api/orders", orderHandler.GetOrders) Maybe add later if needed
	app.Get("/api/orders/:id", orderHandler.GetOrderByID)
	app.Patch("/api/orders/:id", orderHandler.UpdateOrder)
	app.Delete("/api/orders/:id", orderHandler.DeleteOrder)

	app.Get("/api/session/:id/order", orderHandler.GetOrderDetailsBySessionId)

	return app
}
