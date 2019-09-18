package user

// User represents all information of a user
type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Mail     string    `json:"mail"`
	Password string    `json:"password"`
	Clients  *[]string `json:"clients"`
	Admin    bool      `json:"admin"`
}
