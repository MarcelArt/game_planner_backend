package handler

import (
	"game_planner_backend/model"
	"game_planner_backend/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepo repository.IUserRepo
}

func NewUserHandler(userRepo repository.IUserRepo) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// CreateUser is a function to create new User
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body model.UserInput true "Create User"
// @Success 201 {number} integer
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userInput *model.UserInput
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	userInput.Password = string(hashedPassword)

	id, err := h.userRepo.Create(userInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

// GetAllUsers is a function to get all Users saved in system
// @Summary Get all Users
// @Description Get all Users
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param limit query int false "limit"
// @Success 200 {array} model.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	p := c.Query("page", "0")
	l := c.Query("limit", "10")

	page, err := strconv.Atoi(p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	limit, err := strconv.Atoi(l)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	users, err := h.userRepo.Read(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// UpdateUser is a function to update User by id
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param input body model.UserInput true "Update User"
// @Success 200 {object} model.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var userInput *model.UserInput

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	userInput.Password = string(hashedPassword)

	if err := h.userRepo.Update(id, userInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	user, err := h.userRepo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser is a function to delete one User saved in system by id
// @Summary Delete one User
// @Description Delete one User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 500 {string} string
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userRepo.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// GetOneUser is a function to get one User saved in system by id
// @Summary Get one User
// @Description Get one User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 500 {string} string
// @Router /user/{id} [get]
func (h *UserHandler) GetOneUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userRepo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
