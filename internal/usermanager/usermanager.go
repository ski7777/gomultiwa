package usermanager

import (
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

func NewUserManager(c config.ConfigData) *UserManager {
	var um = new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	return um
}
