package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morgansundqvist/service-composable-commerce/internal/application"
	"github.com/morgansundqvist/service-composable-commerce/internal/domain"
)

type OrderHandler struct {
	orderService *application.OrderService
}

func NewOrderHandler(orderService *application.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var input domain.CreateOrderInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order, err := h.orderService.CreateOrder(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := h.orderService.GetOrderById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(order)
}

func (h *OrderHandler) GetOrderDetailsBySessionId(c *fiber.Ctx) error {
	sessionId := c.Params("sessionId")
	orderDetails, err := h.orderService.GetOrderDetailsBySessionId(sessionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(orderDetails)
}

func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var input domain.UpdateOrderInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order, err := h.orderService.UpdateOrder(id, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(order)
}

func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.orderService.DeleteOrder(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
