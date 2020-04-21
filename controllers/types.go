package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mux    *mux.Router
	config string
}

type RouterAPI interface {
	IRouter
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
type IRouter interface {
	SetupRouter(server *http.ServeMux)
}

type AppServer struct {
	Config  string
	Routers []IRouter
	server  *http.ServeMux
}

func NewAppServer() *AppServer {
	return &AppServer{
		server:  http.NewServeMux(),
		Routers: make([]IRouter, 0),
	}
}

func (app *AppServer) AddRouter(router ...IRouter) {
	app.Routers = append(app.Routers, router...)
}

func (app AppServer) SetRoutes() {
	for _, router := range app.Routers {
		router.SetupRouter(app.server)
	}
}

func (app AppServer) StartServer(address string) {

	http.ListenAndServe(address, app.server)
}
