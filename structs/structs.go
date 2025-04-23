package structs


type QueryParams struct {
	Address string `form:"address" binding:"required"`
}

type DomainResult struct {
	Registrar string   `json:"registrar"`
	Record    []string `json:"record"`
}

type TxtResult struct {
	Registrar string   `json:"registrar"`
	Value     []string `json:"value"`
}
