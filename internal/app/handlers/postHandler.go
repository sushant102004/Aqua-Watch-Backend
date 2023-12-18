package handler

import (
	"context"
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

	requiredFields := []string{"UserID", "Date", "Time", "ImageURL", "Description", "Coordinates", "Location"}

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

func (h *PostHandler) HandleGetAllPosts(ctx *fiber.Ctx) error {
	posts, err := h.store.GetAllPosts(context.Background())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "unable to get all posts" + err.Error(),
		})
	}

	return ctx.JSON(map[string]any{
		"posts": posts,
	})
}

func (h *PostHandler) HandleUpdateDamageScore(ctx *fiber.Ctx) error {
	postID := ctx.Query("postID")

	if postID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "postid must be specified in query params",
		})
	}

	err := h.store.IncreaseDamageScore(context.Background(), postID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(http.StatusOK)
}

func (h *PostHandler) HandleSearchPostsViaCity(ctx *fiber.Ctx) error {
	city := ctx.Query("city")

	if city == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "city must be specified in query params",
		})
	}

	posts, err := h.store.SearchPostsVIALocation(context.Background(), city)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if len(posts) == 0 {
		return ctx.SendStatus(http.StatusNoContent)
	}

	return ctx.JSON(map[string][]types.UserPost{
		"data": posts,
	})
}

func (h *PostHandler) HandleGetPostsForMap(ctx *fiber.Ctx) error {

	posts, err := h.store.GetPostsForMap(context.Background())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if len(posts) == 0 {
		return ctx.SendStatus(http.StatusNoContent)
	}

	return ctx.JSON(map[string][]types.UserPostMap{
		"data": posts,
	})
}
