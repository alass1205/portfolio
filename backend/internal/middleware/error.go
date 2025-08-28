package middleware

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError

    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }

    log.Printf("Erreur: %v", err)

    return c.Status(code).JSON(ErrorResponse{
        Success: false,
        Message: "Une erreur s'est produite",
        Error:   err.Error(),
    })
}
