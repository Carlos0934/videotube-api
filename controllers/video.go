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

func (controller *VideoController) getVideo(r *http.Request) (models.Video, error) {
	video := models.Video{}

	err := json.NewDecoder(r.Body).Decode(&video)

	return video, err
}

func (controller *VideoController) getVideofilter(r *http.Request) map[string]interface{} {
	ids := mux.Vars(r)

	return map[string]interface{}{"id": ids["user"], "user_id": ids["video"]}
}
func (controller *VideoController) SetupRouter(server *mux.Router) {
	controller.SetupRouterAPI(server, controller)
}

func (controller *VideoController) Get(w http.ResponseWriter, r *http.Request) {

	video := &models.VideoStorage{}

	filter := controller.getVideofilter(r)
	err := controller.storage.FindOne(filter, video)
	checkErr(err)

	data, err := json.Marshal(video)
	checkErr(err)

	w.Write(data)
}

func (controller *VideoController) Post(w http.ResponseWriter, r *http.Request) {
	video, err := controller.getVideo(r)
	checkErr(err)

	err = controller.storage.Save(&video)
	checkErr(err)

	w.Write(NewResponseMessage("Video created successfully", false))
}

func (controller *VideoController) GetAll(w http.ResponseWriter, r *http.Request) {
	videos := []models.VideoStorage{}
	userId := mux.Vars(r)["user"]
	err := controller.storage.FindByUser(userId, &videos)
	checkErr(err)

	data, err := json.Marshal(videos)
	checkErr(err)

	w.Write(data)

}

func (controller *VideoController) Put(w http.ResponseWriter, r *http.Request) {
	filter := controller.getVideofilter(r)

	video, err := controller.getVideo(r)
	checkErr(err)
	controller.storage.Update(filter, &video)
	w.Write(NewResponseMessage("Video updated successfully", false))

}

func (controller *VideoController) Delete(w http.ResponseWriter, r *http.Request) {
	filter := controller.getVideofilter(r)

	if controller.storage.Delete(filter) {
		w.Write(NewResponseMessage("Video updated deleted", false))
	} else {
		w.Write(NewResponseMessage("Delete video failed ", false))
	}
}
