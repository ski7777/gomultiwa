package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/structs"
)

// ResponseWriter writes respnse based on respnse code, error, json payload (struct), rawpayload ([]byte) and contenttype
func ResponseWriter(w http.ResponseWriter, code int, err error, payload interface{}, rawpayload []byte, rawcontenttype string) {
	if err != nil {
		payload = structs.NewError(err)
	}
	contenttype := ""
	rawdata := make([]byte, 0)
	if payload != nil {
		data, e := json.Marshal(payload)
		if e != nil {
			ResponseWriter(w, 500, e, nil, nil, "")
			return
		}
		rawdata = []byte(data)
		contenttype = "application/json"
	} else if rawpayload != nil {
		rawdata = rawpayload
		contenttype = rawcontenttype
	}
	w.Header().Add("Server", "golang/gomultiwa")
	w.WriteHeader(code)
	if payload == nil && rawpayload == nil {
		return
	}
	if contenttype == "" {
		w.Header().Add("Content-Type", contenttype)
	}
	if _, err := w.Write(rawdata); err != nil {
		log.Println(err)
	}
}
