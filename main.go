package main

import (
	"b11/domainKit/commands"
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
	domain.Use(middlewares.ValidateParams())
	domain.GET("/a", getARecordForAddress)
	domain.GET("/aaaa", getAAAARecordForAddress)
	domain.GET("/ns", getNSRecordForAddress)
	domain.GET("/txt")

	endless.ListenAndServe(":8080", router)
}

func getARecordForAddress(c *gin.Context) {
	returnCode := http.StatusOK
	var results []structs.DomainResult

	params := c.MustGet("params").(structs.QueryParams)

	for _, registrar := range domainList {
		array := commands.ExecuteARecordQuery(registrar, params.Address)

		results = append(results, structs.DomainResult{
			Registrar: registrar,
			Record:    array,
		})
	}
	c.JSON(returnCode, results)
}

func getAAAARecordForAddress(c *gin.Context) {
	returnCode := http.StatusOK
	var results []structs.DomainResult

	params := c.MustGet("params").(structs.QueryParams)

	for _, registrar := range domainList {
		array := commands.ExecuteAAAARecordQuery(registrar, params.Address)

		results = append(results, structs.DomainResult{
			Registrar: registrar,
			Record:    array,
		})
	}
	c.JSON(returnCode, results)
}

func getNSRecordForAddress(c *gin.Context) {
	returnCode := http.StatusOK
	var results []structs.DomainResult

	params := c.MustGet("params").(structs.QueryParams)

	for _, registrar := range domainList {
		array := commands.ExecuteNSRecordQuery(registrar, params.Address)

		results = append(results, structs.DomainResult{
			Registrar: registrar,
			Record:    array,
		})
	}
	c.JSON(returnCode, results)
}