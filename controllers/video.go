package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/carlos0934/videotube/models"
	"github.com/gorilla/mux"
)

type VideoController struct {
	*ControllerAPI
	storage *models.VideoStorage
}

func NewVideoController(conn *sql.DB) *VideoController {
	return &VideoController{
		storage:       models.NewVideoStorage(conn),
		ControllerAPI: NewControllerAPI("/users/{user}/videos", "video"),
	}
}

func (controller *VideoController) getVideo(r *http.Request) (*models.Video, error) {
	video := &models.Video{}

	err := json.NewDecoder(r.Body).Decode(&video)

	return video, err
}

func (controller *VideoController) SetupRouter(server *mux.Router) {
	controller.SetupRouterAPI(server, controller)
}

func (controller *VideoController) Get(w http.ResponseWriter, r *http.Request) {

}

func (controller *VideoController) Post(w http.ResponseWriter, r *http.Request) {

}

func (controller *VideoController) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (controller *VideoController) Put(w http.ResponseWriter, r *http.Request) {

}

func (controller *VideoController) Delete(w http.ResponseWriter, r *http.Request) {

}
