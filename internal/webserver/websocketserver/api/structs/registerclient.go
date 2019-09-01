package structs

type RegisterClientReq struct {
	PNG bool `json:"png"`
}

type RegisterClientRes struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	PNG   string `json:"png"`
}
