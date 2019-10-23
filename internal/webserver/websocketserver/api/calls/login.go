package calls

import (
	"net/http"

	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

// Login executes the login call
func Login(wa gmwi.GoMultiWAInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(structs.LoginReq)
		if err := util.RequestLoader(r, req); err != nil {
			util.ResponseWriter(w, 400, err, nil, nil, "")
			return
		}
		if sess, err := wa.LoginMailPassword(req.Mail, req.Password); err != nil {
			util.ResponseWriter(w, 403, err, nil, nil, "")
		} else {
			res := new(structs.LoginRes)
			res.Session = sess
			util.ResponseWriter(w, 200, nil, structs.NewOKRes(res), nil, "")
		}
	}
}
