package api

import (
	"log"

	"github.com/chensheep/hotel-reservation-backend/db"
	"github.com/chensheep/hotel-reservation-backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.userStore.GetUser(c.Context(), id)
	if err != nil {
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"error": err.Error(),
		// })
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errs := params.Validate(); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	createdUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(createdUser)
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// log.Println(string(c.Body()))

	var values types.UpdateUserParams
	if err := c.BodyParser(&values); err != nil {
		return err
	}

	log.Println(values)

	err := h.userStore.UpdateUser(c.Context(), id, values)
	if err != nil {
		return err
	}

	return nil
}

// handleDeleteUser handles DELETE /api/users/:id
func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
