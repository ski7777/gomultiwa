package scopemanager

import (
	"errors"
	"log"
	"sync"

	"github.com/ski7777/gomsgqueue/pkg/messagequeue"
)

type ScopeManager struct {
	scopes         *[]*Scope
	scopeslock     sync.Mutex
	mq             *messagequeue.MessageQueue
	approvehandler func()
	handlerlock    sync.Mutex
}

func (sm *ScopeManager) SetApproveHandler(f func()) {
	sm.approvehandler = f
}

func (sm *ScopeManager) InScopes(s *Scope) bool {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	return sm.inScopes(s)
}

func (sm *ScopeManager) RequestScope(s *Scope) {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	if !sm.inScopes(s) {
		*sm.scopes = append(*sm.scopes, s)
		go sm.requestScope(s)
	}
}

func (sm *ScopeManager) RequestScopes(s []*Scope) {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	rs := new([]*Scope)
	for _, ss := range s {
		if !sm.inScopes(ss) {
			*sm.scopes = append(*sm.scopes, ss)
			*rs = append(*rs, ss)
		}
	}
	go sm.requestScopes(rs)
}

func (sm *ScopeManager) GetScopes() []*Scope {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	return *sm.scopes
}

func (sm *ScopeManager) GetScopeApproved(s *Scope) (bool, error) {
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			return v.Approved, nil
		}
	}
	return false, errors.New("Scope not found")
}

func (sm *ScopeManager) inScopes(s *Scope) bool {
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			return true
		}
	}
	return false
}

func (sm *ScopeManager) handleApproveSingle(_ string, _ string, pl interface{}, _ bool) {
	s := pl.(*Scope)
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, v := range *sm.scopes {
		if v.EqualsTo(s) {
			*v = *s
		}
	}
	go sm.callApproveHandler()
}

func (sm *ScopeManager) handleApproveMultiple(_ string, _ string, pl interface{}, _ bool) {
	sl := pl.([]interface{})
	sm.scopeslock.Lock()
	defer sm.scopeslock.Unlock()
	for _, si := range sl {
		s := si.(*Scope)
		for _, v := range *sm.scopes {
			if v.EqualsTo(s) {
				*v = *s
			}
		}
	}
	go sm.callApproveHandler()
}

func (sm *ScopeManager) requestScope(s *Scope) {
	if _, e := sm.mq.SendMessage(s, MsgScopeManagerRequestScopeSingle); e != nil {
		log.Fatal(e)
	}
}

func (sm *ScopeManager) requestScopes(s *[]*Scope) {
	if _, e := sm.mq.SendMessage(s, MsgScopeManagerRequestScopeMultiple); e != nil {
		log.Fatal(e)
	}
}

func (sm *ScopeManager) callApproveHandler() {
	sm.handlerlock.Lock()
	defer sm.handlerlock.Unlock()
	if sm.approvehandler != nil {
		sm.approvehandler()
	}
}

func NewScopeManager(mq *messagequeue.MessageQueue) *ScopeManager {
	sm := new(ScopeManager)
	sm.scopes = new([]*Scope)
	sm.mq = mq
	sm.mq.RegisterDataHandler(MsgScopeManagerApproveScopeSingle, sm.handleApproveSingle)
	sm.mq.RegisterDataHandler(MsgScopeManagerApproveScopeMultiple, sm.handleApproveMultiple)
	return sm
}
