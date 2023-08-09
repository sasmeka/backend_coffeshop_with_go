package middleware

import (
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.Responses(401, &config.Result{Message: "Please login"}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.Responses(401, &config.Result{Message: "Invalid header value"}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			pkg.Responses(401, &config.Result{Message: err.Error()}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.Responses(401, &config.Result{Data: "you not have permission to access"}).Send(ctx)
			return
		}

		ctx.Set("userId", check.Id)
		ctx.Next()

	}
}
