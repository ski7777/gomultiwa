package calls

import (
	"encoding/base64"
	"net/http"

	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
	qrcode "github.com/skip2/go-qrcode"
)

func RegisterClient(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(structs.RegisterClientReq)
		if err := util.RequestLoader(w, r, req); err != nil {
			util.ResponseWriter(w, 400, err, nil)
			return
		}
		if u, err := wa.UseSession(req.Session); err != nil {
			util.ResponseWriter(w, 403, err, nil)
		} else {
			qr, id, err := wa.StartRegistration(u.ID)
			if err != nil {
				util.ResponseWriter(w, 500, err, nil)
				return
			}
			res := new(structs.RegisterClientRes)
			res.ID = id
			res.Token = <-qr
			if req.PNG == true {
				png, err := qrcode.Encode(res.Token, qrcode.Medium, -1)
				if err != nil {
					util.ResponseWriter(w, 500, err, nil)
					return
				}
				res.PNG = base64.StdEncoding.EncodeToString(png)
			}
			util.ResponseWriter(w, 200, nil, structs.NewOKRes(res))
		}
	}
}
