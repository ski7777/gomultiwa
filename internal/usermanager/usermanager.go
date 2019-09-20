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

// UserManager represents all users and clients
type UserManager struct {
	Userconfig     *user.Users
	WAClients      *waclient.WAClients
	userconfiglock sync.Mutex
}

// CreateUser creates a new user and returns its id
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

// SetUserPW sets password for given user id
func (um *UserManager) SetUserPW(id string, pw string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	(*um.Userconfig.Users)[n].Password = genPWHash(pw)
	return nil

}

// CheckUserPW checks password for givebn user id
func (um *UserManager) CheckUserPW(id string, pw string) (bool, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return false, err
	}
	return (*um.Userconfig.Users)[n].Password == genPWHash(pw), nil

}

// SetUserName sets name for given user id
func (um *UserManager) SetUserName(id string, name string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	(*um.Userconfig.Users)[n].Name = name
	return nil

}

// SetUserMail sets user mail address for given user id
func (um *UserManager) SetUserMail(id string, mail string) error {
	if len(mail) > 254 || !rxEmail.MatchString(mail) {
		return errors.New("Mailaddress not valid")
	}
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	(*um.Userconfig.Users)[n].Mail = mail
	return nil

}

// SetUserAdmin sets admin status for given user id
func (um *UserManager) SetUserAdmin(id string, admin bool) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	(*um.Userconfig.Users)[n].Admin = admin
	return nil

}

// GetUserAdmin returns admin status for given user id
func (um *UserManager) GetUserAdmin(id string) (bool, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return false, err
	}
	return (*um.Userconfig.Users)[n].Admin, nil

}

// AddUserClient adds client id for given user id
func (um *UserManager) AddUserClient(id string, client string) error {
	if _, ok := um.WAClients.Clients[client]; !ok {
		return errors.New("Client ID invalid")
	}
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	for i := 0; i < len(*(*um.Userconfig.Users)[n].Clients); i++ {
		if (*(*um.Userconfig.Users)[n].Clients)[i] == id {
			return errors.New("Client ID already associated with user")
		}
	}
	*(*um.Userconfig.Users)[n].Clients = append((*(*um.Userconfig.Users)[n].Clients), client)
	return nil
}

// DeleteUserClient removes client id for given user id
func (um *UserManager) DeleteUserClient(id string, client string) error {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return err
	}
	for i := 0; i < len(*(*um.Userconfig.Users)[n].Clients); i++ {
		if (*(*um.Userconfig.Users)[n].Clients)[i] == client {
			*(*um.Userconfig.Users)[n].Clients = append((*(*um.Userconfig.Users)[n].Clients)[:i], (*(*um.Userconfig.Users)[n].Clients)[i+1:]...)
			return nil
		}
	}
	return errors.New("Client ID not associated with user")

}

// GetUserClients returns all client ids for given user id
func (um *UserManager) GetUserClients(id string) (*[]string, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return nil, err
	}
	return (*um.Userconfig.Users)[n].Clients, nil
}

// GetUserByID returns User struct for given user id
func (um *UserManager) GetUserByID(id string) (*user.User, error) {
	um.userconfiglock.Lock()
	defer um.userconfiglock.Unlock()
	var n int
	var err error
	if n, err = um.getUserIndexByID(id); err != nil {
		return nil, err
	}
	return (*um.Userconfig.Users)[n], nil
}

// GetUserIDByMail returns user id for given mail address
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

// CheckUserExists checks if user exists for given user id
func (um *UserManager) CheckUserExists(id string) bool {
	for n := range *um.Userconfig.Users {
		if (*um.Userconfig.Users)[n].ID == id {
			return true
		}
	}
	return false
}

// NewUserManager returns new User
func NewUserManager(c config.Data) *UserManager {
	um := new(UserManager)
	um.Userconfig = c.Userconfig
	um.WAClients = c.WAClients
	return um
}
