package main

import (
	"github.com/carlos0934/videotube/controllers"
	"github.com/carlos0934/videotube/db"
	"github.com/carlos0934/videotube/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := BuildContainer("root@/videotube")

	app.StartServer(":3000")
}

func BuildContainer(uri string) *controllers.AppServer {

	conn := db.MysqlConnector(uri)
	userStorage := models.NewUserStorage(conn)
	userController := controllers.NewUserController(userStorage)
	app := controllers.NewAppServer()
	app.AddRouter(userController)
	app.SetRoutes()

	return app
}
