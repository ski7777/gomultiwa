package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestLoader loads request (JSON) to struct
func RequestLoader(w http.ResponseWriter, r *http.Request, v interface{}) error {
	bodyb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyb, v); err != nil {
		return err
	}
	return nil
}
