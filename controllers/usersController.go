package controllers

import "github.com/gofiber/fiber/v2"

func UserIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UserShow(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UserCreate(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UserUpdate(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UserDelete(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
