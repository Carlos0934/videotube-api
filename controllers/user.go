package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carlos0934/videotube/auth"
	"github.com/carlos0934/videotube/models"
	"github.com/gorilla/mux"
)

type UserController struct {
	*ControllerAPI
	storage   *models.UserStorage
	auth      *auth.UserAuth
	middlware *UserMiddleware
}

func NewUserController(conn *sql.DB) *UserController {
	userAuth := auth.NewUserAuth(conn)
	return &UserController{
		storage:       models.NewUserStorage(conn),
		ControllerAPI: NewControllerAPI("/users", "user"),
		auth:          userAuth,
		middlware:     NewUserMiddleware(userAuth),
	}
}
func (controller *UserController) SetupRouter(server *mux.Router) {

	server.Use(controller.middlware.AuthMiddleware)
	controller.SetupRouterAPI(server, controller)

}
func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)[controller.uri]

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

	user := []models.User{}
	err := controller.storage.Find(&user)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(&user)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(data)
}

func (controller *UserController) Post(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = controller.storage.Save(&user)

	if err != nil {
		fmt.Println(err)
	}
	controller.getToken(w, user)
	w.Write(NewResponseMessage("New User created", false))

}

func (controller *UserController) Put(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	id := mux.Vars(r)[controller.uri]

	filter := map[string]interface{}{"id": id}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = controller.storage.Update(filter, user)

	if err != nil {
		fmt.Println(err)
	}

	controller.getToken(w, user)
	w.Write(NewResponseMessage("User updated", false))
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)[controller.uri]

	filter := map[string]interface{}{"id": id}

	result := controller.storage.Delete(filter)

	if result {
		w.Write(NewResponseMessage("User deleted successfully", false))
	} else {
		w.Write(NewResponseMessage("Failed to  try to delete user", false))
	}
}

func (controller *UserController) getToken(w http.ResponseWriter, user models.User) {

	token := controller.auth.GenerateToken(user)
	w.Header().Add("Authorization", token)
}
