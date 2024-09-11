package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/domain"
	"github.com/sajjadjm/wehub-code-challenge/internal/core/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CSVRecordHandler struct {
	service *services.CSVRecordService
}

func NewCSVRecordHandler(service *services.CSVRecordService) *CSVRecordHandler {
	return &CSVRecordHandler{service: service}
}

func (h *CSVRecordHandler) CreateCSVRecord(c *fiber.Ctx) error {
	record := new(domain.CSVRecord)
	if err := c.BodyParser(record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if errs := ValidateStruct(record); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errs,
		})
	}

	record.ID = primitive.NewObjectID()
	if err := h.service.CreateCSVRecord(record); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(record)
}

func (h *CSVRecordHandler) GetCSVRecordByID(c *fiber.Ctx) error {
	id := c.Params("id")
	record, err := h.service.GetCSVRecordByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.JSON(record)
}

func (h *CSVRecordHandler) UpdateCSVRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	record := new(domain.CSVRecord)

	if err := c.BodyParser(record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if errs := ValidateStruct(record); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errs,
		})
	}

	updatedRecord, err := h.service.UpdateCSVRecord(id, record)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedRecord)
}

func (h *CSVRecordHandler) DeleteCSVRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteCSVRecord(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *CSVRecordHandler) GetAllCSVRecords(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)    // Default to page 1
	limit := c.QueryInt("limit", 10) // Default to limit 10

	records, total, err := h.service.GetAllCSVRecords(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data":  records,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
