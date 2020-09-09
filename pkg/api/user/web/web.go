package web

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"

	"github.com/hi019/fiber-boilerplate/ent"
	"github.com/hi019/fiber-boilerplate/pkg/api/user"
)

type HTTP struct {
	svc user.Service
	validator *validator.Validate
	session *session.Session
}

func NewHttp(svc user.Service, f *fiber.App, v *validator.Validate, s *session.Session) {
	h := &HTTP{svc, v, s}

	f.Post("/signup", h.create)
	f.Post("/login", h.login)
}


func (h *HTTP) create(c *fiber.Ctx) error {
	// Parse request
	req := &struct{
		Email string `validate:"required,email"`
		Password string `validate:"required"`
	}{}
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error parsing JSON",
		}
	}
	if err := h.validator.Struct(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Validation error",
		}
	}

	// Create user
	e, err := h.svc.Create(req.Email, req.Password, c.Context())
	if err != nil {
		switch err.(type) {
		default:
			return err
		case *ent.ConstraintError:
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: "User already exists",
			}
		}
	}

	// Send response
	res := struct{
		Email string
		ID int
	}{e.Email, e.ID}
	return c.JSON(res)
}

func (h *HTTP) login(c *fiber.Ctx) error {
	// Parse request
	req := &struct{
		Email string `validate:"required,email"`
		Password string `validate:"required"`
		Remember string
	}{}
	if err := c.BodyParser(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Error parsing JSON",
		}
	}
	if err := h.validator.Struct(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Validation error",
		}
	}

	// Login user
	e, err := h.svc.Login(req.Email, req.Password, c.Context())
	// Check for not found error
	if err != nil {
		switch err.(type) {
		default:
			return err
		case *ent.NotFoundError:
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: "User not found",
			}
		}
	}

	// Save user id in session
	sess := h.session.Get(c)
	defer sess.Save()

	sess.Set("id", e.ID)


	// Send response
	res := struct{
		Email string
		ID int
	}{e.Email, e.ID}
	return c.JSON(res)
}