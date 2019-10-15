package extensions

import (
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type ExtensionManager struct {
	ws *websocketserver.WSServer
	um *usermanager.UserManager
}

func (em *ExtensionManager) Start() {}

func (em *ExtensionManager) Stop() {}

func NewExtensionManager(ws *websocketserver.WSServer, um *usermanager.UserManager) *ExtensionManager {
	em := new(ExtensionManager)
	return em
}
