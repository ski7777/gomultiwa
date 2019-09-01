package calls

import (
	"encoding/base64"
	"net/http"

	"github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
	qrcode "github.com/skip2/go-qrcode"
)

func RegisterClient(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(structs.RegisterClientReq)
		util.RequestLoader(w, r, req)
		qr, id, err := wa.StartRegistration("")
		if err != nil {
			util.ResponseWriter(w, 500, err, nil)
			return
		}
		var res = new(structs.RegisterClientRes)
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
