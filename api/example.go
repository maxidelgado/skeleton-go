package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/maxidelgado/skeleton-go/domain/example"
	"github.com/maxidelgado/toolkit-go/pkg/logger"
	"github.com/maxidelgado/toolkit-go/pkg/router"
	"go.uber.org/zap"
)

func NewExampleHandler(svc example.Service) router.Handler {
	return exampleHandler{svc: svc}
}

type exampleHandler struct {
	svc example.Service
}

func (h exampleHandler) RegisterRoutes(app *fiber.App) {
	examples := app.Group("/examples")
	{
		examples.Get("/:id", h.getExample)
		examples.Get("/:id/error", h.getExampleError)
	}
}

func (h exampleHandler) getExample(ctx *fiber.Ctx) error {
	log := logger.WithContext(ctx.Context())
	id := ctx.Params("id")
	result, err := h.svc.Get(ctx.Context(), id)
	if err != nil {
		log.Error("exampleHandler.getExample()", zap.Error(err))
		return ctx.JSON(fiber.NewError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(result)
}

func (h exampleHandler) getExampleError(ctx *fiber.Ctx) error {
	apiErr := fiber.NewError(http.StatusInternalServerError, "error response")
	return ctx.Status(apiErr.Code).JSON(apiErr)
}
