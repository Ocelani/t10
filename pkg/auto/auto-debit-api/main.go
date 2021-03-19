package main

import (
	"log"
	"os"

	"github.com/Ocelani/t10/pkg/auto"

	"github.com/gofiber/fiber/v2"
)

var MONGO_URL = os.Getenv("MONGO_URL")

type API struct {
	App     *fiber.App
	Service auto.Service
}

// main initializes the service.
func main() {
	a := &API{
		App:     fiber.New(),
		Service: NewMongoService(MONGO_URL, "autos"),
	}
	defer a.App.Shutdown()

	a.Router()

	log.Fatal(a.App.Listen(":4000"))
}
