package gmwi

import (
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type GoMultiWAInterface interface {
	GetClients() *waclient.WAClients
	StartRegistration(user string) (chan string, string, error)
	LoginMailPassword(user string, password string) (string, error)
	UseSession(sess string) (*user.User, error)
	Stop()
	SaveConfig() error
	GetUserManager() *usermanager.UserManager
}
