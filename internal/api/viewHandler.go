package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/morgansundqvist/service-composable-commerce/internal/application"
	"github.com/morgansundqvist/service-composable-commerce/internal/ports"
)

type ViewHandler struct {
	ProductService *application.ProductService
	OrderService   *application.OrderService
	logger         ports.Logger
}

func NewViewHandler(productService *application.ProductService, orderService *application.OrderService, logger ports.Logger) *ViewHandler {
	return &ViewHandler{
		ProductService: productService,
		OrderService:   orderService,
		logger:         logger,
	}
}

func formatPrice(price int) string {
	priceStr := strconv.Itoa(price)

	if len(priceStr) <= 2 {
		return "0." + priceStr
	}

	return priceStr[:len(priceStr)-2] + "." + priceStr[len(priceStr)-2:]
}

func (h *ViewHandler) HomePage(c *fiber.Ctx) error {
	h.logger.Info("Home page accessed", nil)

	productGroupsWithProducts, err := h.ProductService.GetProductGroupsWithProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to load product groups with products")
	}

	return c.Render("index", fiber.Map{
		"Title":         "Welcome to the Commerce Platform",
		"ProductGroups": productGroupsWithProducts.ProductGroups,
		"FormatPrice":   formatPrice,
	})
}
