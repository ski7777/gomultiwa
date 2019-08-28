package gmwi

import "github.com/ski7777/gomultiwa/internal/waclient"

type GoMultiWAInterface interface {
	GetClients() *waclient.WAClients
}
