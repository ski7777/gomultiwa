package websocketserver

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ski7777/gomultiwa/pkg/extensions/structs"

	"github.com/ski7777/gomsgqueue/pkg/messagequeue"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	gmwi "github.com/ski7777/gomultiwa/internal/gomultiwa/interface"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/calls"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver/api/util"
)

// WSServer represents a http.Server, the gomultiwa instance, a websocket upgrader and router
type WSServer struct {
	server   *http.Server
	wa       gmwi.GoMultiWAInterface
	upgrader *websocket.Upgrader
	router   *mux.Router
}

// WSServerConfig represents the initial config for WSServer
type WSServerConfig struct {
	HTTPServerConfig
	WA gmwi.GoMultiWAInterface
}

// NewWSServer returns new WSServer
func NewWSServer(config *WSServerConfig) *WSServer {
	s := new(WSServer)
	s.wa = config.WA
	s.router = mux.NewRouter()
	//_, filename, _, _ := runtime.Caller(0)
	//webdir := path.Join(path.Dir(filename), "../../../web")
	webbox := packr.New("web", "../../../web")
	//staticdir := path.Join(webdir, "static")
	staticbox := packr.New("web/static", "../../../web/static")
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

// Start stars the server
func (ws *WSServer) Start() error {
	return ws.server.ListenAndServe()
}

func (ws *WSServer) HandleExtensionFunc(method string, path string, mq *messagequeue.MessageQueue) {
	log.Println("Registering External Route: " + "Path: " + path + " ; Method: " + method)
	ws.router.HandleFunc(
		path, func(w http.ResponseWriter, r *http.Request) {
			d, e := ioutil.ReadAll(r.Body)
			if _, e := mq.SendMessageAwaitingResponse(
				&structs.HTTPRequest{Data: d, Error: e, ContentType: r.Header.Get("Content-Type")},
				structs.MsgHTTPRequest,
				func(r interface{}, t string) {
					w.Header().Add("Server", "golang/gomultiwa")
					if t == structs.MsgHTTPResponse {
						rd := r.(*structs.HTTPResponse)
						w.WriteHeader(rd.Code)
						w.Header().Add("Content-Type", rd.ContentType)
						_, _ = w.Write(rd.Data)
					} else {
						w.WriteHeader(502)
						_, _ = w.Write(nil)
					}
				},
				func() {
					w.WriteHeader(504)
					w.Header().Add("Server", "golang/gomultiwa")
					_, _ = w.Write(nil)
				}); e != nil {
				w.WriteHeader(504)
				w.Header().Add("Server", "golang/gomultiwa")
					w.Write(nil)
			}
		}).Methods(method)
				})
}

func registerStaticFile(router *mux.Router, box *packr.Box, name string) {
	router.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		data, err := box.Find(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write(data); err != nil {
				log.Println(err)
			}
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
	case "clients":
		return calls.Clients(ws.wa)
	default:
		log.Fatal(errors.New("API NOT FOUND"))
	}
	return nil
}

func notfound(w http.ResponseWriter, _ *http.Request) {
	util.ResponseWriter(w, 404, errors.New("Not Found"), nil, nil, "")
}
