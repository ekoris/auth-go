package permission

import (
	"goravel/app/http/requests/permission"
	"goravel/app/models"

	"math"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PermissionController struct {
	//Dependent services
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		//Inject services
	}
}

func (r *PermissionController) Index(ctx http.Context) http.Response {

	type PaginatedResponse struct {
		Total       int64               `json:"total"`
		PerPage     int                 `json:"per_page"`
		CurrentPage int                 `json:"current_page"`
		LastPage    int                 `json:"last_page"`
		From        int                 `json:"from"`
		To          int                 `json:"to"`
		Result      []models.Permission `json:"data"`
	}

	var permissions []models.Permission

	if ctx.Request().Input("limit") != "" {
		var total int64
		page := 1
		perPage := 10

		if ctx.Request().Input("page") != "" {
			page, _ = strconv.Atoi(ctx.Request().Input("page"))
		}

		perPage, _ = strconv.Atoi(ctx.Request().Input("limit"))

		facades.Orm().Query().Paginate(page, perPage, &permissions, &total)

		lastPage := int(math.Ceil(float64(total) / float64(perPage)))
		from := (page-1)*perPage + 1
		to := page * perPage
		if to > int(total) {
			to = int(total)
		}

		response := PaginatedResponse{
			Total:       total,
			PerPage:     perPage,
			CurrentPage: page,
			LastPage:    lastPage,
			From:        from,
			To:          to,
			Result:      permissions,
		}
		return ctx.Response().Success().Json(http.Json{
			"message": "successfully",
			"data":    response,
		})
	} else {
		facades.Orm().Query().Get(&permissions)
		response := permissions
		return ctx.Response().Success().Json(http.Json{
			"message": "successfully",
			"data":    response,
		})
	}

}

func (r *PermissionController) Store(ctx http.Context) http.Response {
	var permissionRequest permission.PermissionRequest
	validator, errorValidate := ctx.Request().ValidateRequest(&permissionRequest)
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

	permission := models.Permission{
		Name: ctx.Request().Input("name"),
	}

	if err := facades.Orm().Query().Create(&permission); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"errors": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "successfully",
		"data":    permission,
	})
}

func (r *PermissionController) Update(ctx http.Context) http.Response {
	var permissionRequest permission.PermissionRequest
	validator, errorValidate := ctx.Request().ValidateRequest(&permissionRequest)
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

	permission := models.Permission{
		Name: ctx.Request().Input("name"),
	}

	facades.Orm().Query().Where("id", ctx.Request().Input("id")).Update(&permission)

	var response models.Permission
	facades.Orm().Query().Find(&response, "id=?", ctx.Request().Input("id"))

	return ctx.Response().Success().Json(http.Json{
		"message": "successfully",
		"data":    response,
	})
}

func (r *PermissionController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *PermissionController) Delete(ctx http.Context) http.Response {
	return nil
}
