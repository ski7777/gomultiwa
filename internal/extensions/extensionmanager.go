package extensions

import (
	"log"
	"sync"

	"github.com/google/uuid"

	"github.com/ski7777/gomsgqueue/pkg/interfaces"
	"github.com/ski7777/gomsgqueue/pkg/messagequeue"
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type ExtensionManager struct {
	ws             *websocketserver.WSServer
	um             *usermanager.UserManager
	extensions     map[string]*Extension
	extensionslock sync.Mutex
}

func (em *ExtensionManager) Start() {
	em.extensionslock.Lock()
	defer em.extensionslock.Unlock()
	for ei := range em.extensions {
		em.extensions[ei].start()
	}
}

func (em *ExtensionManager) Stop() {
	wait := new(sync.WaitGroup)
	em.extensionslock.Lock()
	defer em.extensionslock.Unlock()
	wait.Add(len(em.extensions))
	for ei := range em.extensions {
		go func(e *Extension) {
			e.stop()
			wait.Done()
		}(em.extensions[ei])
	}
	wait.Wait()
}

func (em *ExtensionManager) HandleLoadError(e error) {}

func (em *ExtensionManager) AddDedicatedExtension(cmd string) {
	e := NewExtension(em.ws, em.um)
	if err := e.ConnectDedicated(cmd); err != nil {
		em.HandleLoadError(err)
	}
	em.addExtension(e)
}

func (em *ExtensionManager) AddEmbeddedExtension(f func(*messagequeue.MessageQueue) interfaces.Extension) {
	e := NewExtension(em.ws, em.um)
	e.ConnectEmbedded(f)
	em.addExtension(e)
}

func (em *ExtensionManager) addExtension(e *Extension) {
	em.extensionslock.Lock()
	defer em.extensionslock.Unlock()
	if id, err := uuid.NewUUID(); err != nil {
		log.Panic(err)
	} else {
		em.extensions[id.String()] = e
	}
}

func NewExtensionManager(ws *websocketserver.WSServer, um *usermanager.UserManager) *ExtensionManager {
	em := new(ExtensionManager)
	em.extensions = make(map[string]*Extension)
	em.loadExtensions()
	return em
}
