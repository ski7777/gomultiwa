package util

import "strings"

func WidToJid(w string) string {
	return strings.Split(w, "@")[0] + "@s.whatsapp.net"
}
