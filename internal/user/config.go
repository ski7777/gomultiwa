package user

type Users struct {
	Users *[]*User `json:"users"`
}

func NewUsers() *Users {
	u := new(Users)
	u.Users = &[]*User{}
	return u
}
