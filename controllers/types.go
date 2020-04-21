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
	GetOne(condition map[string]string, pointer interface{}) error
	GetAll(codition map[string]string, pointer interface{}) error
	Save(data interface{}) error
	Update(condition map[string]string, data interface{}) error
	Delete(condition map[string]string) bool
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
