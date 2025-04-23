package main

import (
	"b11/domainKit/middlewares"
	"b11/domainKit/structs"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var domainList = []string{"1.1.1.1", "8.8.8.8"}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	domain := v1.Group("/domain")
	domain.Use(handlers.ValidateParams())
	domain.GET("/a", getARecordForAddress)
	domain.GET("/aaaa")
	domain.GET("/ns")
	domain.GET("/txt")

	endless.ListenAndServe(":8080", router)
}

func getARecordForAddress(c *gin.Context) {
	returnCode := http.StatusOK
	var results []structs.DomainResult

	params := c.MustGet("params").(structs.QueryParams)

	for _, registrar := range domainList {
		array := ExecuteARecordQuery(registrar, params.Address)

		results = append(results, structs.DomainResult{
			Registrar: registrar,
			Record:    array,
		})
	}
	c.JSON(returnCode, results)
}
