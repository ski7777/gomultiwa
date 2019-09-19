package handlerhub

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
)

// HandlerHub implements all message handlers
type HandlerHub struct{}

// HandleError handles type error
func (hh *HandlerHub) HandleError(err error, id string) {}

// HandleTextMessage handles type whatsapp.TextMessage
func (hh *HandlerHub) HandleTextMessage(message whatsapp.TextMessage, id string) {}

// HandleImageMessage handles type whatsapp.ImageMessage
func (hh *HandlerHub) HandleImageMessage(message whatsapp.ImageMessage, id string) {}

// HandleDocumentMessage handles type whatsapp.DocumentMessage
func (hh *HandlerHub) HandleDocumentMessage(message whatsapp.DocumentMessage, id string) {}

// HandleVideoMessage handles type erwhatsapp.VideoMessageror
func (hh *HandlerHub) HandleVideoMessage(message whatsapp.VideoMessage, id string) {}

// HandleAudioMessage handles type whatsapp.AudioMessage
func (hh *HandlerHub) HandleAudioMessage(message whatsapp.AudioMessage, id string) {}

// HandleLocationMessage handles type whatsapp.LocationMessage
func (hh *HandlerHub) HandleLocationMessage(message whatsapp.LocationMessage, id string) {}

// HandleLiveLocationMessage handles type whatsapp.LiveLocationMessage
func (hh *HandlerHub) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage, id string) {}

// HandleJsonMessage handles type string (JSON)
func (hh *HandlerHub) HandleJsonMessage(message string, id string) {}
