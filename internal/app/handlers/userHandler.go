package handler

import (
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/aqua-watch-backend/internal/app/store"
	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
)

type UserHandler struct {
	store store.UserStore
}

func NewUserHandler(store store.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleCreateUser(ctx *fiber.Ctx) error {
	var params types.User

	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "invalid request body",
		})
	}

	requiredFields := []string{"Name", "Email", "Location", "Language", "ProfilePicture", "PhoneNumber"}

	for _, field := range requiredFields {
		value := reflect.ValueOf(params).FieldByName(field).String()
		if value == "" {
			return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": field + " is required",
			})
		}
	}

	resp, err := h.store.CreateUser(ctx.Context(), &params)

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

func (h *UserHandler) HandleLoginUSer(ctx *fiber.Ctx) error {
	email := ctx.Query("email")

	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "please provide email in query",
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
