package user

// Users represents a list of all users
type Users struct {
	Users *[]*User `json:"users"`
}

// NewUsers returns a new Users struct
func NewUsers() *Users {
	u := new(Users)
	u.Users = &[]*User{}
	return u
}
