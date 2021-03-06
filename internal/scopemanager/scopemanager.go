package scopemanager

import (
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/ski7777/gomsgqueue/pkg/messagequeue"
	pkg "github.com/ski7777/gomultiwa/pkg/scopemanager"
)

type ScopeManager struct {
	scopes         *[]*pkg.Scope
	scopeslock     sync.Mutex
	mq             *messagequeue.MessageQueue
	requesthandler func()
	handlerlock    sync.Mutex
}

func (sm *ScopeManager) SetRequestHandler(f func()) {
	sm.requesthandler = f
}

func (sm *ScopeManager) InScopes(s *pkg.Scope) bool {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	return sm.inScopes(s)
}

func (sm *ScopeManager) ApproveScope(s *pkg.Scope) error {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			v.Approved = true
			sm.approveScope(s)
			return nil
		}
	}
	return errors.New("Scope not found")
}

func (sm *ScopeManager) ApproveScopes(s []*pkg.Scope) error {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	var nf int
	var found bool
	rs := new([]*pkg.Scope)
	for _, ss := range s {
		found = false
		for _, v := range *sm.scopes {
			if v.EqualsTo(ss) {
				v.Approved = true
				*rs = append(*rs, ss)
				found = true
			}
		}
		if !found {
			nf++
		}
	}
	sm.approveScopes(rs)
	if nf > 0 {
		return errors.New(strconv.Itoa(nf) + " Scope(s) not found")
	}
	return nil
}

func (sm *ScopeManager) GetScopes() []*pkg.Scope {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	return *sm.scopes
}

func (sm *ScopeManager) GetScopeApproved(s *pkg.Scope) (bool, error) {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			return v.Approved, nil
		}
	}
	return false, errors.New("Scope not found")
}

func (sm *ScopeManager) inScopes(s *pkg.Scope) bool {
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			return true
		}
	}
	return false
}

func (sm *ScopeManager) handleRequestSingle(_ string, _ string, pl interface{}, _ bool) {
	s := pl.(*pkg.Scope)
	if !sm.inScopes(s) {
		sm.scopeslock.Lock()
		defer sm.scopeslock.Unlock()
		*sm.scopes = append(*sm.scopes, s)
		go sm.callRequestHandler()
	}
}

func (sm *ScopeManager) handleRequestMultiple(_ string, _ string, pl interface{}, _ bool) {
	sl := pl.([]interface{})
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, si := range sl {
		s := si.(*pkg.Scope)
		if !sm.inScopes(s) {
			*sm.scopes = append(*sm.scopes, s)
		}
	}
	go sm.requesthandler()
}

func (sm *ScopeManager) approveScope(s *pkg.Scope) {
	if _, e := sm.mq.SendMessage(s, pkg.MsgScopeManagerApproveScopeSingle); e != nil {
		log.Fatal(e)
	}
}

func (sm *ScopeManager) approveScopes(s *[]*pkg.Scope) {
	if _, e := sm.mq.SendMessage(s, pkg.MsgScopeManagerApproveScopeMultiple); e != nil {
		log.Fatal(e)
	}
}

func (sm *ScopeManager) callRequestHandler() {
	sm.handlerlock.Lock()
	defer sm.handlerlock.Unlock()
	if sm.requesthandler != nil {
		sm.requesthandler()
	}
}

func NewScopeManager(mq *messagequeue.MessageQueue) *ScopeManager {
	sm := new(ScopeManager)
	sm.scopes = new([]*pkg.Scope)
	sm.mq = mq
	sm.mq.RegisterDataHandler(pkg.MsgScopeManagerRequestScopeSingle, sm.handleRequestSingle)
	sm.mq.RegisterDataHandler(pkg.MsgScopeManagerRequestScopeMultiple, sm.handleRequestMultiple)
	return sm
}
