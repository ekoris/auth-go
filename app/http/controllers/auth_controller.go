package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/http/requests/Auth"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	//Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

func (r *AuthController) Register(ctx http.Context) http.Response {

	var registerRequest requests.RegisterRequest
	validator, errorValidate := ctx.Request().ValidateRequest(&registerRequest)

	if errorValidate != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
			"error": errorValidate,
		})
	}

	if validator != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
			"message": validator.All(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ctx.Request().Input("password")), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"error": err.Error(),
		})
	}

	user := models.User{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: string(hashedPassword),
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"errors sd": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Data Register successfully",
		"data":    user,
	})
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	var loginRequest Auth.LoginRequest
	validator, errorValidate := ctx.Request().ValidateRequest(&loginRequest)

	if errorValidate != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
			"error": errorValidate,
		})
	}

	if validator != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
			"message": validator.All(),
		})
	}

	// Retrieve user from database
	var user models.User
	if err := facades.Orm().Query().Where("email", ctx.Request().Input("email")).First(&user); err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{
			"error": "Invalid credentials",
		})
	}

	// Compare provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ctx.Request().Input("password"))); err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{
			"error": "Invalid credentials",
		})
	}

	// Generate token
	token, err := facades.Auth(ctx).LoginUsingID(user.ID)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Data Login successfully",
		"data": map[string]string{
			"user":  user.Name,
			"email": user.Email,
			"token": token,
		},
	})
}
