package handler

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/aqua-watch-backend/internal/app/store"
	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
)

type NGOHandler struct {
	store store.NGOStore
}

func NewNGOHandler(store store.NGOStore) *NGOHandler {
	return &NGOHandler{
		store: store,
	}
}

func (h *NGOHandler) HandleSignUp(ctx *fiber.Ctx) error {
	var params types.NGO

	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "invalid request body",
		})
	}

	requiredFields := []string{"Name", "Email", "PhoneNumber", "Description", "Location", "ImageUrl"}

	for _, field := range requiredFields {
		value := reflect.ValueOf(params).FieldByName(field)
		if value.IsZero() {
			return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": field + " is required",
			})
		}
	}

	err := h.store.SignUp(context.Background(), params)
	if err != nil {
		return ctx.JSON(map[string]string{
			"error": "error: " + err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(map[string]string{
		"message": "account created successfully",
	})
}

func (h *NGOHandler) HandleLogin(ctx *fiber.Ctx) error {
	email := ctx.Query("email")

	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "email is required in query parameters",
		})
	}

	resp, err := h.store.Login(ctx.Context(), email)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "success",
		"user":    resp,
	})
}
