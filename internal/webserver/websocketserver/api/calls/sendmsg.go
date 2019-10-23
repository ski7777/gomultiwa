package calls

import (
	"errors"
	"net/http"

	whatsapp "github.com/Rhymen/go-whatsapp"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

// SendMsg executes the sendmsg call
func SendMsg(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(structs.SendmsgReq)
		if err := util.RequestLoader(r, req); err != nil {
			util.ResponseWriter(w, 400, err, nil, nil, "")
			return
		}
		if u, err := wa.UseSession(req.Session); err != nil {
			util.ResponseWriter(w, 403, err, nil, nil, "")
		} else {
			found := false
			for _, c := range *u.Clients {
				if c == req.ID {
					found = true
				}
			}
			if !found {
				util.ResponseWriter(w, 404, errors.New("WA Client ID not found"), nil, nil, "")
				return
			}
			if _, err := wa.GetClients().Clients[req.ID].WAClient.WA.Send(whatsapp.TextMessage{
				Info: whatsapp.MessageInfo{
					RemoteJid: req.JID,
				},
				Text: req.MSG,
			}); err != nil {
				util.ResponseWriter(w, 500, err, nil, nil, "")
				return
			}
			util.ResponseWriter(w, 200, nil, structs.NewOKRes(nil), nil, "")
		}
	}
}
