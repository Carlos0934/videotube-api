package controllers

import (
	"crypto/ecdsa"
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

func NewUserController(conn *sql.DB, privateKey *ecdsa.PrivateKey) *UserController {
	userAuth := auth.NewUserAuth(conn, privateKey)
	return &UserController{
		storage:       models.NewUserStorage(conn),
		ControllerAPI: NewControllerAPI("/users", "user"),
		auth:          userAuth,
		middlware:     NewUserMiddleware(userAuth),
	}
}
func (controller *UserController) SetupRouter(server *mux.Router) {

	server.Use(controller.middlware.AuthMiddleware)
	server.HandleFunc("/auth", controller.ValidateUser)

	controller.SetupRouterAPI(server, controller)

}

func (controller *UserController) addToken(w http.ResponseWriter, user models.User) {

	token := controller.auth.GenerateToken(user)
	w.Header().Add("Authorization", token)
}
func (controller *UserController) getUser(r *http.Request) models.User {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	return user
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)[controller.uri]

	filter := map[string]interface{}{"id": id}
	user := models.User{}
	err := controller.storage.FindOne(filter, &user)

	data, err := json.Marshal(&user)

	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write(NewResponseMessage("Error to try get user", true))
	}

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

	if err == nil {
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write(NewResponseMessage("Error to try get users", true))
	}
}

func (controller *UserController) Post(w http.ResponseWriter, r *http.Request) {

	user := controller.getUser(r)
	err := controller.storage.Save(&user)

	if err != nil {
		fmt.Println(err)
	}
	controller.addToken(w, user)

	if err == nil {
		w.WriteHeader(200)
		w.Write(NewResponseMessage("user created successfully", false))
	} else {
		w.WriteHeader(400)
		w.Write(NewResponseMessage("Error to try create user", true))
	}

}

func (controller *UserController) Put(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)[controller.uri]

	filter := map[string]interface{}{"id": id}

	user := controller.getUser(r)
	err := controller.storage.Update(filter, user)

	if err != nil {
		fmt.Println(err)
	}

	controller.addToken(w, user)
	if err == nil {
		w.WriteHeader(200)
		w.Write(NewResponseMessage("User updated successfully", false))
	} else {
		w.WriteHeader(400)
		w.Write(NewResponseMessage("Error to try update user", true))
	}
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)[controller.uri]

	filter := map[string]interface{}{"id": id}

	result := controller.storage.Delete(filter)

	if result {
		w.Write(NewResponseMessage("User deleted successfully", false))
	} else {
		w.Write(NewResponseMessage("Failed to  try to delete user", true))
	}
}

func (controller *UserController) ValidateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	user := controller.getUser(r)
	if user.Password == "" {
		w.Write(NewResponseMessage("Password not provided", false))
		return
	}
	if controller.auth.VerifyPassword(user) {
		controller.addToken(w, user)
		w.Write(NewResponseMessage("Token successfully created ", false))
		return
	}

	w.Write(NewResponseMessage("Invalid user", false))
}
