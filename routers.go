package routers

import (
	"net/http"
	"timesheet-api/controllers"
	"timesheet-api/middlewares"
	"timesheet-api/validators"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// Middleware configuration
	authMiddleWare := middlewares.NewAuthMiddleware()

	// Controller configuration
	healthzController := controllers.NewHealthzController()
	exampleController := controllers.NewExampleController()
	leaveController := controllers.NewLeaveController()

	// Router definition
	app.Get("/healthz", healthzController.GetHealthz)
	app.Get("/example", exampleController.GetExample)
	app.Get("/example/error", exampleController.ErrorExample)
	//test api
	app.Get("/hello", controllers.GetHello)
	app.Post("/create-request", authMiddleWare.Auth, controllers.CreateRequest)
	http.HandleFunc("/slack/intueractionsr", controllers.HandleInteractionEvent)
	leave := app.Group("/leaves", authMiddleWare.Auth)
	leave.Get("/", leaveController.GetLeaves)
	leave.Post("/", validators.CreateLeaveValidator, leaveController.CreateLeave)

}
