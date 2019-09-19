package structs

// ErrorRes represents Error response struct
type ErrorRes struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

// NewError returns new ErrorRes
func NewError(e error) *ErrorRes {
	es := new(ErrorRes)
	es.Status = "Error"
	es.Reason = e.Error()
	return es
}
