package main

import (
	"fmt"

	"github.com/Ocelani/t10/pkg/auto"
	"github.com/gofiber/fiber/v2"
)

// Router defines the URL routing for Router.
func (a *API) Router() {
	a.App.Post("/auto-debit/add", a.add())
	a.App.Get("/auto-debit/all", a.getAll())
	a.App.Get("/auto-debit/:find", a.getOne()) // Name or ID
	a.App.Get("/auto-debit", a.queryStatus())
	a.App.Put("/auto-debit/:id/approve", a.updateStatus("approved"))
	a.App.Put("/auto-debit/:id/reject", a.updateStatus("rejected"))
	a.App.Delete("/auto-debit/:id", a.remove())
}

// add through API service request.
func (a *API) add() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody auto.Entity
		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(400).SendString(
				fmt.Sprint(err, reqBody))
		}
		found, _ := a.Service.FindOneWithName(reqBody.Name)
		if !found.ID.IsZero() {
			return c.Status(200).JSON(found)
		}
		// reqBody.Status = "pending"
		resp, err := a.Service.Insert(reqBody)
		if err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err))
		}

		return c.Status(201).JSON(resp)
	}
}

func (a *API) queryStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		q := c.Query("status")
		result, err := a.Service.FindAllWithStatus(q)
		if err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err, q))
		}

		return c.Status(200).JSON(result)
	}
}

// getAll through API client request.
func (a *API) getAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := a.Service.FindAll()
		if err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err))
		}

		return c.Status(200).JSON(result)
	}
}

// getOneWithID through API client request.
func (a *API) getOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			err    error
			result auto.Entity
			params = c.Params("find")
		)
		if len(params) == 24 { // id
			result, err = a.Service.FindOneWithID(params)
		} else { // name
			result, err = a.Service.FindOneWithName(params)
		}
		if err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err))
		}

		return c.Status(200).JSON(result)
	}
}

func (a *API) updateStatus(status string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := a.Service.FindOneWithID(id)
		if err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err))
		}
		result.Status = status

		return c.Status(200).JSON(result)
	}
}

// remove through API client request.
func (a *API) remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := a.Service.Remove(id); err != nil {
			return c.Status(422).SendString(
				fmt.Sprint(err))
		}

		return c.Status(200).SendString(id)
	}
}
