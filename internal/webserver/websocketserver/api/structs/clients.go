package structs

type ClientsReq struct {
	Session string `json:"sess"`
}

type ClientsRes struct {
	Clients map[string]string `json:"clients"`
}
