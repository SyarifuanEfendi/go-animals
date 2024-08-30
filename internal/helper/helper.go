package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseInt(value string) (int, error) {
    result, err := strconv.Atoi(value)
    if err != nil {
        return 0, err
    }
    return result, nil
}

func ParseIntFromParam(c *fiber.Ctx, paramName string) (int, error) {
    value := c.Params(paramName)
    if value == "" {
        return 0, fiber.NewError(fiber.StatusBadRequest, "Missing or invalid parameter")
    }
    return ParseInt(value)
}