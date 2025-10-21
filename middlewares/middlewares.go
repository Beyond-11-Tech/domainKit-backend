package middlewares

import (
	"b11/domainKit/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// this function pre-validates the query params for the domain endpoints, either aborting the execution if they are missing
// or setting the params in a "params" variable and passing to the handler
func ValidateParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params structs.QueryParams

		err := c.ShouldBindQuery(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}

		c.Set("params", params)

		c.Next()
	}
}

// this function is to validate the bearer key against the presented key in the authorisation.
// it will call in the bearer token and compare it against the env list of allowed keys.
// it will pass on if valid, and 401 if invalid.
func ValidateAuth(validKey []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authData := ctx.Request.Header.Values("Authorization")

		for _, item := range validKey {
			token := strings.Split(authData[0], " ")[1]
			if token == item {
				ctx.Next()
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorised"})
		ctx.Abort()

	}
}
