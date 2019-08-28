package structs

type Error struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func NewError(e error) *Error {
	es := new(Error)
	es.Status = "Error"
	es.Reason = e.Error()
	return es
}
