package structs

type ErrorRes struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func NewError(e error) *ErrorRes {
	es := new(ErrorRes)
	es.Status = "Error"
	es.Reason = e.Error()
	return es
}
