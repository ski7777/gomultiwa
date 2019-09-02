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

func NewUserManager(c config.ConfigData) *UserManager {
	var um = new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	return um
}
