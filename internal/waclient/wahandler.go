package waclient

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
)

type WAHandler struct {
	id *uuid.UUID
	hh *handlerhub.HandlerHub
}

func (wah *WAHandler) HandleError(err error) {
	wah.hh.HandleError(err, wah.id)
}

func (wah *WAHandler) HandleTextMessage(message whatsapp.TextMessage) {
	wah.hh.HandleTextMessage(message, wah.id)
}

func (wah *WAHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	wah.hh.HandleImageMessage(message, wah.id)
}

func (wah *WAHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	wah.hh.HandleDocumentMessage(message, wah.id)
}

func (wah *WAHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	wah.hh.HandleVideoMessage(message, wah.id)
}

func (wah *WAHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	wah.hh.HandleAudioMessage(message, wah.id)
}

func (wah *WAHandler) HandleLocationMessage(message whatsapp.LocationMessage) {
	wah.hh.HandleLocationMessage(message, wah.id)
}
func (wah *WAHandler) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage) {
	wah.hh.HandleLiveLocationMessage(message, wah.id)
}

func (wah *WAHandler) HandleJsonMessage(message string) {
	wah.hh.HandleJsonMessage(message, wah.id)
}
