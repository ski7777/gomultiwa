package calls

import (
	"errors"
	"net/http"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

func SendMsg(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(structs.SendmsgReq)
		if err := util.RequestLoader(w, r, req); err != nil {
			util.ResponseWriter(w, 400, err, nil)
			return
		}
		wacc, ok := wa.GetClients().Clients[req.ID]
		if !ok {
			util.ResponseWriter(w, 404, errors.New("WA Client ID not found!"), nil)
			return
		}
		wac := wacc.WAClient.WA
		if _, err := wac.Send(whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: req.JID,
			},
			Text: req.MSG,
		}); err != nil {
			util.ResponseWriter(w, 500, err, nil)
			return
		}
		util.ResponseWriter(w, 200, nil, structs.NewOKRes(nil))
	}
}
