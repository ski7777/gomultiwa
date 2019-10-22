package extensions

import (
	"github.com/ski7777/goextensioniser/pkg/extensioniser"
	"github.com/ski7777/gomsgqueue/pkg/interfaces"
	"github.com/ski7777/gomsgqueue/pkg/messagequeue"
	sm "github.com/ski7777/gomultiwa/internal/scopemanager"
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type Extension struct { // implements "github.com/ski7777/gomsgqueue/pkg/interfaces".Master
	ext interfaces.Extension
	sm  *sm.ScopeManager
	ws  *websocketserver.WSServer
	um  *usermanager.UserManager
	mq  *messagequeue.MessageQueue
	mqm *messagequeue.Master
}

func (e *Extension) ConnectEmbedded(f func(*messagequeue.MessageQueue) interfaces.Extension) {
	mqe := messagequeue.NewExtension()
	mqe.ConnectToMaster(e.mqm)
	e.ext = f(mqe.GetMessageQueue())
}

func (e *Extension) ConnectDedicated(cmd string) error {
	cm, err := extensioniser.NewDedicatedExtension("go run $GOPATH/src/github.com/ski7777/gomsgqueue/examples/extension/dedicated/*.go")
	if err != nil {
		return err
	}
	cm.ConnectToMaster(e.mqm)
	return nil
}

func (e *Extension) start() {
	e.mq.Run()
}

func (e *Extension) stop() {
	e.mq.Stop()
}

func (e *Extension) handleScopeRequest() {}

func NewExtension(ws *websocketserver.WSServer, um *usermanager.UserManager) *Extension {
	e := new(Extension)
	e.ws = ws
	e.um = um
	e.mqm = messagequeue.NewMaster()
	e.mq = e.mqm.GetMessageQueue()
	e.sm = sm.NewScopeManager(e.mq)
	e.sm.SetRequestHandler(e.handleScopeRequest)
	return e
}
