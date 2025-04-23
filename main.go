package main

import (
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type domainResult struct {
	Registrar string   `json:"registrar"`
	Record    []string `json:"record"`
}

var domainList = []string{"1.1.1.1", "8.8.8.8"}

func main() {
	router := gin.Default()
	router.GET("/domain", getRecordForDomain)

	endless.ListenAndServe(":8080", router)
}

func getRecordForDomain(c *gin.Context) {
	type queryParams struct {
		Address string `form:"address" binding:"required"`
	}
	returnCode := http.StatusOK
	var results []domainResult

	var params queryParams

	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, registrar := range domainList {
		buf := ExecuteARecordQuery(registrar, params.Address)

		results = append(results, domainResult{
			Registrar: registrar,
			Record:    buf,
		})
	}
	c.JSON(returnCode, results)

}
