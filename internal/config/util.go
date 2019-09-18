package config

import (
	"crypto/sha512"
)

func getConfigHash(config []byte) string {
	return string(sha512.New().Sum(config))
}
