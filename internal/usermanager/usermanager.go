package usermanager

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type UserManager struct {
	Userconfig *user.Users
	WAClients  *waclient.WAClients
}

func (um *UserManager) CreateUser(name string, mail string) (string, error) {
	for _, u := range *um.Userconfig.Users {
		if u.Mail == mail {
			return "", errors.New("Mailaddress already registered")
		}
	}
	i, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := i.String()
	var u = new(user.User)
	u.ID = id
	u.Name = name
	u.Mail = mail
	*um.Userconfig.Users = append(*um.Userconfig.Users, u)
	return id, nil
}

func (um *UserManager) SetUserPW(id string, pw string) error {
	for n := range *um.Userconfig.Users {
		if (*um.Userconfig.Users)[n].ID == id {
			x := genPWHash(pw)
			(*um.Userconfig.Users)[n].Password = x
			return nil

		}
	}
	return errors.New("User ID not found")
}

func NewUserManager(c config.ConfigData) *UserManager {
	var um = new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	//um.Userconfig.Users = append(um.Userconfig.Users, &user.User{[]string{"a", "b"}, []user.UserPermisson{}, false})
	return um
}
