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

func (app AppServer) SetRoutes() {
	for _, router := range app.Routers {
		router.SetupRouter(app.server)
	}
}
