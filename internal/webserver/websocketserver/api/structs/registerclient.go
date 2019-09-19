package structs

// RegisterClientReq represents the requets struct for registerclient call
type RegisterClientReq struct {
	PNG     bool   `json:"png"`
	PNGRAW  bool   `json:"pngraw"`
	Session string `json:"sess"`
}

// RegisterClientRes represents the response struct for registerclient call
type RegisterClientRes struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	PNG   string `json:"png"`
}
