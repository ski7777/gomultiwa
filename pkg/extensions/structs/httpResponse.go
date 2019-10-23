package structs

const MsgHTTPResponse = "http/response"

type HTTPResponse struct {
	Data        []byte `json:"Data"`
	Code        int    `json:"Code"`
	ContentType string `json:"ContentType"`
}
