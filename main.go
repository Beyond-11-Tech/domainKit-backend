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
	router.GET("/domain", getRecordForDomain)

	router.Run(":8080")
}

func getRecordForDomain(c *gin.Context) {
	domain, present := c.GetQuery("domain")
	returnCode := http.StatusOK

	var results []domainResult

	if !present {
		c.String(http.StatusBadRequest, "please include a 'domain' query param")
		return
	}

	for _, registrar := range domainList {
		buf, err := executeDigQuery(registrar, domain)
		record := strings.Fields(buf.String())

		if err != nil {
			results = append(results, domainResult{
				Registrar: "error",
				Record:    []string{"Domain: " + registrar, err.Error()},
			})
			returnCode = http.StatusInternalServerError
			continue
		}

		results = append(results, domainResult{
			Registrar: registrar,
			Record:    record,
		})
	}
	c.JSON(returnCode, results)

}
