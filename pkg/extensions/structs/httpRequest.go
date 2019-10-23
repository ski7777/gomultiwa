package structs

const MsgHTTPRequest="http/request"

type HTTPRequest struct {
	Data  []byte `json:"Data"`
	Error error  `json:"error"`	
	ContentType string `json:"ContentType"`
}
