package structs


type QueryParams struct {
	Address string `form:"address" binding:"required"`
}

type DomainResult struct {
	Registrar string   `json:"registrar" example:"1.1.1.1"`
	Record    []string `json:"record" example:"21.11.25.4,1.54.23.6"`
}

type TxtResult struct {
	Registrar string   `json:"registrar"`
	Value     []string `json:"value"`
}
