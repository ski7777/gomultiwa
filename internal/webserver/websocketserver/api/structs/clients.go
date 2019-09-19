package structs

// ClientsReq represents the requets struct for clients call
type ClientsReq struct {
	Session string `json:"sess"`
}

// ClientsRes represents the resopnse struct for clients call
type ClientsRes struct {
	Clients map[string]string `json:"clients"`
}
