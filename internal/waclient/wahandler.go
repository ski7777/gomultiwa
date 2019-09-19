package waclient

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
)

// WAHandler implemets the whatsapp.Handler interface and sends all data to handlerhub
type WAHandler struct {
	id string
	hh *handlerhub.HandlerHub
}

// SetID sets the waclient id
func (wah *WAHandler) SetID(id string) {
	wah.id = id
}

// HandleError handles type error
func (wah *WAHandler) HandleError(err error) {
	wah.hh.HandleError(err, wah.id)
}

// HandleTextMessage handles type whatsapp.TextMessage
func (wah *WAHandler) HandleTextMessage(message whatsapp.TextMessage) {
	wah.hh.HandleTextMessage(message, wah.id)
}

// HandleImageMessage handles type whatsapp.ImageMessage
func (wah *WAHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	wah.hh.HandleImageMessage(message, wah.id)
}

// HandleDocumentMessage handles type whatsapp.DocumentMessage
func (wah *WAHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	wah.hh.HandleDocumentMessage(message, wah.id)
}

// HandleVideoMessage handles type erwhatsapp.VideoMessageror
func (wah *WAHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	wah.hh.HandleVideoMessage(message, wah.id)
}

// HandleAudioMessage handles type whatsapp.AudioMessage
func (wah *WAHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	wah.hh.HandleAudioMessage(message, wah.id)
}

// HandleLocationMessage handles type whatsapp.LocationMessage
func (wah *WAHandler) HandleLocationMessage(message whatsapp.LocationMessage) {
	wah.hh.HandleLocationMessage(message, wah.id)
}

// HandleLiveLocationMessage handles type whatsapp.LiveLocationMessage
func (wah *WAHandler) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage) {
	wah.hh.HandleLiveLocationMessage(message, wah.id)
}

// HandleJsonMessage handles type string (JSON)
func (wah *WAHandler) HandleJsonMessage(message string) {
	wah.hh.HandleJsonMessage(message, wah.id)
}
