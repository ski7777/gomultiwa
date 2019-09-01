package structs

type OKRes struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}

func NewOKRes(p interface{}) *OKRes {
	es := new(OKRes)
	es.Status = "OK"
	es.Payload = p
	return es
}
