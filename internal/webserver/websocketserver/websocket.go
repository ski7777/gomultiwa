package websocketserver

import (
	"errors"
	"log"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/calls"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

type WSServer struct {
	server   *http.Server
	wa       gmwi.GoMultiWAInterface
	upgrader *websocket.Upgrader
	router   *mux.Router
}

type WSServerConfig struct {
	HTTPServerConfig
	WA gmwi.GoMultiWAInterface
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
	s.router.HandleFunc("/api/v1/sendmsg", s.apihandler("sendmsg")).Methods("POST")
	s.router.HandleFunc("/api/v1/registerclient", s.apihandler("registerclient")).Methods("POST")
	s.router.HandleFunc("/api/v1/login", s.apihandler("login")).Methods("POST")
	s.router.HandleFunc("/api/v1/clients", s.apihandler("clients")).Methods("POST")
	s.router.NotFoundHandler = s.router.NewRoute().HandlerFunc(notfound).GetHandler()
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

func (ws *WSServer) apihandler(call string) func(http.ResponseWriter, *http.Request) {
	switch call {
	case "sendmsg":
		return calls.SendMsg(ws.wa)
	case "registerclient":
		return calls.RegisterClient(ws.wa)
	case "login":
		return calls.Login(ws.wa)
	default:
		log.Fatal(errors.New("API NOT FOUND"))
	}
	return nil
}

func notfound(w http.ResponseWriter, _ *http.Request) {
	util.ResponseWriter(w, 404, errors.New("Not Found!"), nil)
}
