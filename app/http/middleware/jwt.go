package middleware

import (
	"errors"

	"github.com/goravel/framework/auth"
	httpcontract "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt() httpcontract.Middleware {
	return func(ctx httpcontract.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			response := map[string]string{"error": "Authorization token is missing"}
			ctx.Request().AbortWithStatusJson(401, response)
			return
		}

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth(ctx).Refresh()
				if err != nil {
					response := map[string]string{"error": "Token refresh failed"}
					ctx.Request().AbortWithStatusJson(401, response)
					return
				}

				token = "Bearer " + token
			} else {

				response := map[string]string{"error": "Invalid token"}
				ctx.Request().AbortWithStatusJson(401, response)
				return
			}
		}

		// You can get User in DB and set it to ctx

		//var user models.User
		//if err := facades.Auth().User(ctx, &user); err != nil {
		//	ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		//  return
		//}
		//ctx.WithValue("user", user)

		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}
}
