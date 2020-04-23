package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carlos0934/videotube/models"
	"github.com/gorilla/mux"
)

type UserController struct {
	Controller
	storage models.UserStorage
}

func NewUserController() *UserController {

	return &UserController{}
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["user"]
	filter := map[string]interface{}{"id": id}
	user := models.User{}
	err := controller.storage.FindOne(filter, &user)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(&user)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)

}

func (controller *UserController) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Post(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Put(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {

}
