package middlewares

import (
	"b11/domainKit/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//this function pre-validates the query params for the domain endpoints, either aborting the execution if they are missing
//or setting the params in a "params" variable and passing to the handler
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
