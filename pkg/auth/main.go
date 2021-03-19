package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	App   *fiber.App
	Store map[string]jwt.MapClaims
}

// main initializes the service.
func main() {
	a := &API{
		App: fiber.New(),
		Store: map[string]jwt.MapClaims{
			"jwt-secret": jwt.MapClaims{"authorized": true}},
	}
	defer a.App.Shutdown()

	a.App.Post("/auth/add/:pass", a.register())
	a.App.Get("/auth", a.authenticate())

	log.Fatal(a.App.Listen(":4000"))
}

func (a *API) register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		pass := c.Params("pass")
		fmt.Println(pass)
		val := jwt.MapClaims{
			"username": pass,
			"admin":    true,
			"exp":      time.Now().Add(time.Minute * 15).Unix(),
		}
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, val)
		key, err := at.SignedString([]byte(pass))
		if err != nil {
			fmt.Println(key, err)
			return c.Status(400).SendString(
				fmt.Sprint(err))
		}
		a.Store[key] = val

		return c.Status(201).SendString(key)
	}
}

func (a *API) authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get("X-Auth")

		for k := range a.Store {
			if key == k {
				c.Set("X-Auth", key)
				fmt.Println(key)
				return c.Status(200).SendString(
					"Authenticated")
			}
		}

		return c.Status(401).SendString(
			"Unauthorized")
	}
}
