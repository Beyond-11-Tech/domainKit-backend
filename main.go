package main

import (
	"b11/domainKit/commands"
	"b11/domainKit/middlewares"
	"b11/domainKit/structs"
	"flag"
	"net/http"

	docs "b11/domainKit/docs"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type systemFlag struct {
	webKey string
	appKey string
}

var systemFlags systemFlag
var domainList = []string{"1.1.1.1", "8.8.8.8"}

func init() {
	flag.StringVar(&systemFlags.webKey, "webKey", "", "web authentication key, used to give basic auth permissions to the API")
	flag.StringVar(&systemFlags.appKey, "appKey", "", "app authentication key, used to give basic auth permissions to the API")
	flag.Parse()

	if systemFlags.webKey == "" || systemFlags.appKey == "" {
		panic("missing required flags, please add 'appKey' and 'webKey' and run again")
	}
	//go:generate swag init
	//go:generate ./...

}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.Title="Domainkit API"
	docs.SwaggerInfo.Description="The backend API for the B11 domainkit system"
	docs.SwaggerInfo.Version="1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "domainkit.choccobear.tech"
	docs.SwaggerInfo.Schemes = []string{"https"}

	v1 := router.Group("/v1")
	v1.GET("/health", healthCheck)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	domain := v1.Group("/domain", gin.BasicAuth(gin.Accounts{"web": systemFlags.webKey, "app": systemFlags.appKey}))
	domain.Use(middlewares.ValidateParams())
	domain.GET("/a", getARecordForAddress)
	domain.GET("/aaaa", getAAAARecordForAddress)
	domain.GET("/ns", getNSRecordForAddress)
	domain.GET("/srv", notAvailableYet)
	domain.GET("/txt", notAvailableYet)

	endless.ListenAndServe(":8080", router)
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func notAvailableYet(c *gin.Context) {
	c.String(http.StatusTeapot, "im still brewing")
}

// @Tags Domains
// @Router /domain/a [get]
// @Summary get domain A records for given domain
// @Param domain query string true "valid (sub)domain to query" minlength(4)
// @Produce json
// @Success 200 {array} structs.DomainResult
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

// @Tags Domains
// @Router /domain/aaaa [get]
// @Summary get domain AAAA (IPv6) records for given domain
// @Param domain query string true "valid (sub)domain to query" minlength(4)
// @Produce json
// @Success 200 {array} structs.DomainResult
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

// @Tags Domains
// @Router /domain/ns [get]
// @Summary get NS records for given domain
// @Param domain query string true "valid (sub)domain to query" minlength(4)
// @Produce json
// @Success 200 {array} structs.DomainResult
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
