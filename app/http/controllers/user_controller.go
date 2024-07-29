package controllers

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Fetch(ctx http.Context) http.Response {
	var user models.User

	// Populate the user struct
	err := facades.Auth(ctx).User(&user)
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(map[string]string{
			"error": "User not authenticated",
		})
	}

	// Verify the populated user
	fmt.Println(user)

	// Return user information
	return ctx.Response().Success().Json(map[string]interface{}{
		"Hello": user,
	})
}
