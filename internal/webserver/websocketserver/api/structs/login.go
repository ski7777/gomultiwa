package structs

type LoginReq struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type LoginRes struct {
	Session string `json:"sess"`
}
