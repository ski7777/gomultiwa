package main

import (
	"github.com/ski7777/gomultiwa/internal/gomultiwa"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

func main() {
	var gmw = new(gomultiwa.GoMultiWA)
	var wsc = new(websocketserver.WSServerConfig)
	wsc.Host = "0.0.0.0"
	wsc.Port = 8888
	wsc.WA = gmw
	var ws = websocketserver.NewWSServer(wsc)
	ws.Start()
	<-make(chan int, 1)
}
