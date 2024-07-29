package permission

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type PermissionRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *PermissionRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *PermissionRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required",
	}
}

func (r *PermissionRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
