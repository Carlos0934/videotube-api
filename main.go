package main

import (
	"github.com/carlos0934/videotube/controllers"
	"github.com/carlos0934/videotube/db"
	"github.com/carlos0934/videotube/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := BuildContainer()

	app.StartServer(":3000")
}

func BuildContainer() *controllers.AppServer {

	conn := db.MysqlConnector("root@/videotube")
	userStorage := models.NewUserStorage(conn)
	userController := controllers.NewUserController(userStorage)
	app := controllers.NewAppServer()
	app.AddRouter(userController)

	return app
}
