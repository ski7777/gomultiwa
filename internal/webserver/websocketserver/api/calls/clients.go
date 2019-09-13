package calls

import (
	"net/http"

	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

func Clients(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(structs.ClientsReq)
		if err := util.RequestLoader(w, r, req); err != nil {
			util.ResponseWriter(w, 400, err, nil)
			return
		}
		if u, err := wa.UseSession(req.Session); err != nil {
			util.ResponseWriter(w, 403, err, nil)
		} else {
			res := new(structs.ClientsRes)
			res.Clients = make(map[string]string)
			clients := wa.GetClients().Clients
			for _, i := range *u.Clients {
				res.Clients[i] = clients[i].Session.Wid
			}
			util.ResponseWriter(w, 200, nil, structs.NewOKRes(res))
		}
	}
}
