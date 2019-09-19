package structs

// LoginReq represents the requets struct for login call
type LoginReq struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

// LoginRes represents the response struct for login call
type LoginRes struct {
	Session string `json:"sess"`
}
