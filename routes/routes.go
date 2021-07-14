package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

// *fiber.App App denotes the Fiber application.
// Setup routing the application
func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	/*
		app.Post("/api/login", controllers.Login)

		app.Use(middlewares.IsAuthenticated)

		app.Get("/api/user", controllers.User)
		app.Post("/api/logout", controllers.Logout)


		app.Get("/api/users", controllers.AllUsers)
		// changes roles to role
		app.Post("/api/roles", controllers.CreateRole)
		// changes roles to role
		app.Get("/api/roles/:id",controllers.GetRole)
		// changes roles to role
		app.Put("/api/roles/:id",controllers.UpdateRole)
		// changes roles to role
		app.Delete("/api/roles/:id", controllers.DeleteRole)

		app.Get("/api/permissions", controllers.AllPermissions)

		app.Get("/api/products", controllers.AllProducts)
		// changes products to product
		app.Post("/api/products", controllers.CreateProduct)
		// changes products to product
		app.Get("/api/products/:id", controllers.GetProduct)
		// changes products to product
		app.Put("/api/products/:id",controllers.UpdateProduct)
		// changes products to product
		app.Delete("/api/products/:id", controllers.DeleteProduct)

		app.Post("/api/upload", controllers.Upload)
		// Static will create a file server serving static files
		app.Static("/api/uploads", "./uploads")


		app.Get("/api/orders", controllers.AllOrders)
		app.Post("/api/export",controllers.Export)
		app.Get("/api/chart", controllers.Chart)
	*/
}
