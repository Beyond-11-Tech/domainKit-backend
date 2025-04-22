package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type domainResult struct {
	Registrar string   `json:"registrar"`
	Record    []string `json:"record"`
}

var domainList = []string{"1.1.1.1", "8.8.8.8"}

func main() {
	router := gin.Default()
	router.GET("/google", getDnsForGoogle)
	router.GET("/domain", getRecordForDomain)

	router.Run(":8080")
}

func getDnsForGoogle(c *gin.Context) {
	var results []domainResult

	for _, domain := range domainList {
		buf := executeDigQuery(domain, "google.com")

		results = append(results, domainResult{
			Registrar: domain,
			Record:    strings.Split(buf.String(), "\n"),
		})
	}
	c.JSON(http.StatusOK, results)
}

func getRecordForDomain(c *gin.Context) {
	domain, present := c.GetQuery("domain")
	var results []domainResult

	if !present {
		c.String(http.StatusBadRequest, "please include a 'domain' query param")
		return
	}

	for _, registrar := range domainList {
		buf := executeDigQuery(registrar, domain)
		record := strings.Fields(buf.String())

		results = append(results, domainResult{
			Registrar: registrar,
			Record:    record,
		})
	}
	c.JSON(http.StatusOK, results)

}
