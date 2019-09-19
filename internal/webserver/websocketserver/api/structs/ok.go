package structs

// OKRes represents the ok response struct
type OKRes struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}

// NewOKRes returns new OKRes
func NewOKRes(p interface{}) *OKRes {
	es := new(OKRes)
	es.Status = "OK"
	es.Payload = p
	return es
}
