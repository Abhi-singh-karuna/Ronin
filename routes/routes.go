package routes

import (
	"github.com/amshashankk/authentication"
	"github.com/gofiber/fiber/v2"
)


//--------------Created Routing funtion for SignUp & login-------------//
func Route(app fiber.Router) {
	app.Post("/login", authentication.Login)
	app.Post("/signup", authentication.SignUp)
}