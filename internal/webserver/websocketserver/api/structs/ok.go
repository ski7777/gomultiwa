package structs

type OK struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}

func NewOK(p interface{}) *OK {
	es := new(OK)
	es.Status = "OK"
	es.Payload = p
	return es
}
