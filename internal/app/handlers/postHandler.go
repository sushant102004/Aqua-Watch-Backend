package handler

import (
	"net/http"

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
	var params *types.UserPost

	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.store.InsertPost(ctx.Context(), params)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]string{
		"message": resp,
	})
}
