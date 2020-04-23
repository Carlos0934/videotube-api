package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type IControllerAPI interface {
	IController
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
type IController interface {
	SetupRouter(server *mux.Router)
}

type AppServer struct {
	Config  string
	Routers []IController
	server  *mux.Router
}

func NewAppServer() *AppServer {
	return &AppServer{
		server:  mux.NewRouter(),
		Routers: make([]IController, 0),
	}
}

func (app *AppServer) AddRouter(router ...IController) {
	app.Routers = append(app.Routers, router...)
}

func (app *AppServer) SetRoutes() {
	for _, router := range app.Routers {
		router.SetupRouter(app.server)
	}
}

func (app *AppServer) StartServer(address string) {
	app.SetRoutes()
	http.ListenAndServe(address, app.server)
}

type ControllerAPI struct {
	uri string
	url string
}

type Message struct {
	Content string `json:"message"`
	IsError bool   `json:"error"`
}

func NewResponseMessage(content string, err bool) []byte {
	message := &Message{
		Content: content,
		IsError: err,
	}

	data, _ := json.Marshal(message)

	return data
}

func NewControllerAPI(url, uri string) *ControllerAPI {
	return &ControllerAPI{
		url: url,
		uri: uri,
	}
}
func (controller *ControllerAPI) getUriPath() string {
	return fmt.Sprintf("/{%v}", controller.uri)
}
func (internal *ControllerAPI) SetupRouterAPI(server *mux.Router, controller IControllerAPI) {
	route := server.PathPrefix(internal.url).Subrouter()
	route.HandleFunc("", controller.GetAll).Methods("GET")
	route.HandleFunc("", controller.Post).Methods("POST")

	uri := internal.getUriPath()
	route.HandleFunc(uri, controller.Get).Methods("GET")
	route.HandleFunc(uri, controller.Put).Methods("PUT")
	route.HandleFunc(uri, controller.Delete).Methods("DELETE")

}
