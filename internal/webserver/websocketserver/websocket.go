package websocketserver

import (
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	wac "github.com/ski7777/gomultiwa/internal/waclient"
)

type WSServer struct {
	server   *http.Server
	wa       *wac.WAClient
	upgrader *websocket.Upgrader
	router   *mux.Router
}

type WSServerConfig struct {
	HTTPServerConfig
	WA *wac.WAClient
}

func NewWSServer(config *WSServerConfig) *WSServer {
	s := new(WSServer)
	s.wa = config.WA
	s.router = mux.NewRouter()
	_, filename, _, _ := runtime.Caller(0)
	var webdir = path.Join(path.Dir(filename), "../../../web")
	var webbox = packr.New(webdir, webdir)
	var staticdir = path.Join(webdir, "static")
	var staticbox = packr.New(staticdir, staticdir)
	registerStaticFile(s.router, webbox, "index.html")
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(staticbox)))
	s.server = &http.Server{
		Addr:         config.GetAddr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}
	s.upgrader = new(websocket.Upgrader)
	return s
}

func (s *WSServer) Start() error {
	return s.server.ListenAndServe()
}

func registerStaticFile(router *mux.Router, box *packr.Box, name string) {
	router.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		data, err := box.Find(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	})
}

