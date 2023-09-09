package handler

import (
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/aqua-watch-backend/internal/app/store"
	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
)

type PostHandler struct {
	store store.PostStore
}

func NewPostHandler(store store.PostStore) *PostHandler {
	return &PostHandler{store: store}
}

func (h *PostHandler) HandleInsertPost(ctx *fiber.Ctx) error {
	var params types.UserPost

	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "invalid request body",
		})
	}

	requiredFields := []string{"UserID", "Date", "Time", "ImageURL", "Description", "DamageScore", "Coordinates"}

	for _, field := range requiredFields {
		value := reflect.ValueOf(params).FieldByName(field)
		if value.IsZero() {
			return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": field + " is required",
			})
		}
	}

	resp, err := h.store.InsertPost(ctx.Context(), &params)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]string{
		"message": resp,
	})
}
