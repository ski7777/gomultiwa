package usermanager

import (
	"crypto/sha512"
	"encoding/base64"
)

func genPWHash(pw string) string {
	return base64.StdEncoding.EncodeToString(sha512.New().Sum([]byte(pw)))
}
