package handlerhub

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
)

type HandlerHub struct{}

func (hh *HandlerHub) HandleError(err error, id *uuid.UUID)                                          {}
func (hh *HandlerHub) HandleTextMessage(message whatsapp.TextMessage, id *uuid.UUID)                 {}
func (hh *HandlerHub) HandleImageMessage(message whatsapp.ImageMessage, id *uuid.UUID)               {}
func (hh *HandlerHub) HandleDocumentMessage(message whatsapp.DocumentMessage, id *uuid.UUID)         {}
func (hh *HandlerHub) HandleVideoMessage(message whatsapp.VideoMessage, id *uuid.UUID)               {}
func (hh *HandlerHub) HandleAudioMessage(message whatsapp.AudioMessage, id *uuid.UUID)               {}
func (hh *HandlerHub) HandleLocationMessage(message whatsapp.LocationMessage, id *uuid.UUID)         {}
func (hh *HandlerHub) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage, id *uuid.UUID) {}
func (hh *HandlerHub) HandleJsonMessage(message string, id *uuid.UUID)                               {}
