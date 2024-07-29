package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RegisterRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *RegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RegisterRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required",
		"email":    "required",
		"password": "required",
	}
}

func (r *RegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RegisterRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
