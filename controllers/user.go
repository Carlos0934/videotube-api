package controllers

import (
	"net/http"

	"github.com/carlos0934/videotube/models"
)

type UserController struct {
	Controller
	storage models.UserStorage
}

func NewUserController() *UserController {

	return &UserController{}
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Post(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Put(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {

}
