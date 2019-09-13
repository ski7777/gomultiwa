package usermanager

import (
	"errors"
	"regexp"
	"sync"

	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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
	u := new(user.User)
	u.ID = id
	u.Name = name
	u.Mail = mail
	u.Clients = &[]string{}
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
func (um *UserManager) SetUserName(id string, name string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		(*um.Userconfig.Users)[n].Name = name
		return nil
	}
}
func (um *UserManager) SetUserMail(id string, mail string) error {
	if len(mail) > 254 || !rxEmail.MatchString(mail) {
		return errors.New("Mailaddress not valid")
	}
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		(*um.Userconfig.Users)[n].Mail = mail
		return nil
	}
}
func (um *UserManager) SetUserAdmin(id string, admin bool) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		(*um.Userconfig.Users)[n].Admin = admin
		return nil
	}
}
func (um *UserManager) GetUserAdmin(id string) (bool, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return false, err
	} else {
		return (*um.Userconfig.Users)[n].Admin, nil
	}
}
func (um *UserManager) AddUserClient(id string, client string) error {
	if _, ok := um.WAClients.Clients[client]; !ok {
		return errors.New("Client ID invalid")
	}
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		for i := 0; i < len(*(*um.Userconfig.Users)[n].Clients); i++ {
			if (*(*um.Userconfig.Users)[n].Clients)[i] == id {
				return errors.New("Client ID already associated with user")
			}
		}
		*(*um.Userconfig.Users)[n].Clients = append((*(*um.Userconfig.Users)[n].Clients), client)
	}
	return nil
}

func (um *UserManager) DeleteUserClient(id string, client string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return err
	} else {
		for i := 0; i < len(*(*um.Userconfig.Users)[n].Clients); i++ {
			if (*(*um.Userconfig.Users)[n].Clients)[i] == client {
				*(*um.Userconfig.Users)[n].Clients = append((*(*um.Userconfig.Users)[n].Clients)[:i], (*(*um.Userconfig.Users)[n].Clients)[i+1:]...)
				return nil
			}
		}
		return errors.New("Client ID not associated with user")
	}
}

func (um *UserManager) GetUserClients(id string) (*[]string, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return nil, err
	} else {
		return (*um.Userconfig.Users)[n].Clients, nil
	}
}

func (um *UserManager) GetUserByID(id string) (*user.User, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	if n, err := um.getUserIndexByID(id); err != nil {
		return nil, err
	} else {
		return (*um.Userconfig.Users)[n], nil
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

func (um *UserManager) CheckUserExists(id string) bool {
	for n := range *um.Userconfig.Users {
		if (*um.Userconfig.Users)[n].ID == id {
			return true
		}
	}
	return false
}

func NewUserManager(c config.ConfigData) *UserManager {
	um := new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	return um
}
