package usermanager

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type UserManager struct {
	Userconfig     *user.Users
	WAClients      *waclient.WAClients
	userconfiglock sync.Mutex
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
	u.Clients = &[]string{}
	u.Permissions = &[]*user.UserPermisson{}
	*um.Userconfig.Users = append(*um.Userconfig.Users, u)
	return id, nil
}

func (um *UserManager) SetUserPW(id string, pw string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		(*um.Userconfig.Users)[n].Password = genPWHash(pw)
		return nil
	}
}

func (um *UserManager) CheckUserPW(id string, pw string) (bool, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return false, err
	} else {
		return (*um.Userconfig.Users)[n].Password == genPWHash(pw), nil
	}
}

func (um *UserManager) GetUserIDByMail(mail string) (string, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	for n := range *um.Userconfig.Users {
		if (*um.Userconfig.Users)[n].Mail == mail {
			return (*um.Userconfig.Users)[n].ID, nil
		}
	}
	return "", errors.New("Mailaddress not found")
}

func (um *UserManager) getUserIndexByID(id string) (int, error) {
	for n := range *um.Userconfig.Users {
		if (*um.Userconfig.Users)[n].ID == id {
			return n, nil

		}
	}
	return -1, errors.New("User ID not found")
}

func NewUserManager(c config.ConfigData) *UserManager {
	var um = new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	//um.Userconfig.Users = append(um.Userconfig.Users, &user.User{[]string{"a", "b"}, []user.UserPermisson{}, false})
	return um
}
