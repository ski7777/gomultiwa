package handlerhub

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
)

type HandlerHub struct{}

func (hh *HandlerHub) HandleError(err error, id string)                                          {}
func (hh *HandlerHub) HandleTextMessage(message whatsapp.TextMessage, id string)                 {}
func (hh *HandlerHub) HandleImageMessage(message whatsapp.ImageMessage, id string)               {}
func (hh *HandlerHub) HandleDocumentMessage(message whatsapp.DocumentMessage, id string)         {}
func (hh *HandlerHub) HandleVideoMessage(message whatsapp.VideoMessage, id string)               {}
func (hh *HandlerHub) HandleAudioMessage(message whatsapp.AudioMessage, id string)               {}
func (hh *HandlerHub) HandleLocationMessage(message whatsapp.LocationMessage, id string)         {}
func (hh *HandlerHub) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage, id string) {}
func (hh *HandlerHub) HandleJsonMessage(message string, id string)                               {}
