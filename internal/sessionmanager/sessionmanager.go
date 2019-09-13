package sessionmanager

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/session"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/usermanager"
)

const (
	timeout = int64(180 * 24 * time.Hour)
)

type SessionManager struct {
	um          *usermanager.UserManager
	sessions    map[string]*session.Session
	sessionlock sync.Mutex
}

func (sm *SessionManager) NewSession(user string) (string, error) {
	if sm.um.CheckUserExists(user) {
		i, _ := uuid.NewRandom()
		id := i.String()
		sm.sessionlock.Lock()
		defer sm.sessionlock.Unlock()
		sm.sessions[id] = session.NewSession(timeout, user)
		return id, nil
	}
	return "", errors.New("User does not exist")
}

func (sm *SessionManager) Cleanup() {
	sm.sessionlock.Lock()
	defer sm.sessionlock.Unlock()
	for n, s := range sm.sessions {
		if !s.IsValid() {
			delete(sm.sessions, n)
		}
	}
}

func (sm *SessionManager) UseSession(sess string) (*user.User, error) {
	sm.sessionlock.Lock()
	defer sm.sessionlock.Unlock()
	if s, ok := sm.sessions[sess]; ok {
		if err := s.Use(); err != nil {
			return nil, err
		}
		if u, err := sm.um.GetUserByID(s.GetUserID()); err != nil {
			return nil, err
		} else {
			return u, nil
		}
	} else {
		return nil, errors.New("Session ID invalid")
	}

}

func NewSessionManager(um *usermanager.UserManager) *SessionManager {
	sm := new(SessionManager)
	sm.um = um
	sm.sessions = make(map[string]*session.Session)
	return sm
}
