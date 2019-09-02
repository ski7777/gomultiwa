package user

import "log"

type Users struct {
	Users []*User `json:"users"`
}

func NewUsers() *Users {
	var u = new(Users)
	u.Users = []*User{}
	return u
}
