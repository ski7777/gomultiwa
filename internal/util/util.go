package util

import "strings"

// WidToJid converts a Wid to a Jid
func WidToJid(w string) string {
	return strings.Split(w, "@")[0] + "@s.whatsapp.net"
}
