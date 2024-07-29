package routes

import (
	"goravel/app/http/controllers"
	"goravel/app/http/controllers/permission"

	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {

	facades.Route().Prefix("api/v1").Group(func(router route.Router) {
		router.Prefix("/auth").Group(func(router route.Router) {
			authController := controllers.NewAuthController()
			router.Post("/register", authController.Register)
			router.Post("/login", authController.Login)
		})

		router.Middleware(middleware.Jwt()).Group(func(router route.Router) {
			router.Prefix("user").Group(func(router route.Router) {
				userController := controllers.NewUserController()
				router.Get("/fetch", userController.Fetch)
			})

			router.Prefix("permission").Group(func(router route.Router) {
				permissionController := permission.NewPermissionController()
				router.Get("/list", permissionController.Index)
				router.Post("/store", permissionController.Store)
				router.Put("/{id}/update", permissionController.Update)
				router.Delete("/{id}/delete", permissionController.Update)
				router.Get("/{id}/detail", permissionController.Show)
			})
		})

	})

}
