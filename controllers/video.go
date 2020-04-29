package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	ID, err := strconv.Atoi(mux.Vars(r)["user"])
	if err != nil {
		return video, err
	}
	video.UserID = ID
	err = json.NewDecoder(r.Body).Decode(&video)

	return video, err
}

func (controller *VideoController) getVideofilter(r *http.Request) map[string]interface{} {
	ids := mux.Vars(r)

	return map[string]interface{}{"id": ids["video"], "user_id": ids["user"]}
}

func (controller *VideoController) SetupRouter(server *mux.Router) {
	server.HandleFunc("static/{content}/{filename}", controller.GetVideo)
	controller.SetupRouterAPI(server, controller)
}

func (controller *VideoController) getContent(content string, key string, r *http.Request) string {
	err := r.ParseMultipartForm(30 << 20)
	if err != nil {
		fmt.Println(err)
	}

	file, header, err := r.FormFile(key)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	filename := controller.storage.SerializeContent(data, header.Filename, content)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("%v/%v/%v/%v", r.Host, "static", content, filename)
	return url

}

func (controller *VideoController) getMediaContents(r *http.Request) map[string]string {
	videoURL := controller.getContent(models.ContentVideo, "video", r)
	coverURL := controller.getContent(models.ContentCover, "cover", r)

	return map[string]string{"video": videoURL, "cover": coverURL}
}
func (controller *VideoController) Get(w http.ResponseWriter, r *http.Request) {

	video := &models.Video{}

	filter := controller.getVideofilter(r)

	err := controller.storage.FindOne(filter, video)
	checkErr(err, w)

	data, err := json.Marshal(video)
	checkErr(err, w)

	if err == nil {
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write(NewResponseMessage("Error to try get video", true))
	}
}

func (controller *VideoController) Post(w http.ResponseWriter, r *http.Request) {
	video, err := controller.getVideo(r)
	checkErr(err, w)
	contents := controller.getMediaContents(r)

	video.Cover = contents["cover"]
	video.URL = contents["video"]

	err = controller.storage.Save(&video)
	checkErr(err, w)
	if err == nil {
		w.WriteHeader(201)
		w.Write(NewResponseMessage("Video created successfully", false))
	} else {
		w.WriteHeader(400)
		w.Write(NewResponseMessage("Error to try create video", true))
	}

}

func (controller *VideoController) GetAll(w http.ResponseWriter, r *http.Request) {
	videos := []models.Video{}
	userId := mux.Vars(r)["user"]
	err := controller.storage.FindByUser(userId, &videos)
	checkErr(err, w)

	data, err := json.Marshal(&videos)
	checkErr(err, w)

	if err == nil {
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write(NewResponseMessage("Error to try get videos", true))
	}

}

func (controller *VideoController) Put(w http.ResponseWriter, r *http.Request) {
	filter := controller.getVideofilter(r)

	video, err := controller.getVideo(r)
	checkErr(err, w)
	controller.storage.Update(filter, &video)
	if err == nil {
		w.WriteHeader(200)
		w.Write(NewResponseMessage("Video updated successfully", false))
	} else {
		w.WriteHeader(400)
		w.Write(NewResponseMessage("Error to try update video", true))
	}

}

func (controller *VideoController) Delete(w http.ResponseWriter, r *http.Request) {
	filter := controller.getVideofilter(r)

	if controller.storage.Delete(filter) {
		w.WriteHeader(200)
		w.Write(NewResponseMessage("Video deleted succesfully", false))
	} else {
		w.WriteHeader(400)
		w.Write(NewResponseMessage("Delete video failed ", true))
	}
}

func (controller *VideoController) GetVideo(w http.ResponseWriter, r *http.Request) {
	content := mux.Vars(r)["content"]

	filename := mux.Vars(r)["filename"]

	data := controller.storage.DeserializeContent(filename, content)

	w.Write(data)
}
