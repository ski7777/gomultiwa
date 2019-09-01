package util

import (
	"encoding/json"
	"net/http"

	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
)

func ResponseWriter(w http.ResponseWriter, code int, err error, payload interface{}) {
	if err != nil {
		payload = structs.NewError(err)
	}
	data, e := json.Marshal(payload)
	if e != nil {
		ResponseWriter(w, 500, e, "")
	} else {
		w.WriteHeader(code)
		w.Write([]byte(data))
	}
}
