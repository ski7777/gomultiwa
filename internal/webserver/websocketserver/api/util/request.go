package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestLoader(w http.ResponseWriter, r *http.Request, v interface{}) error {
	bodyb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	if err := json.Unmarshal(bodyb, v); err != nil {
		log.Print(err)
		w.WriteHeader(400)
		return err
	}
	return nil
}
