package main

import (
	"github.com/carlos0934/videotube/auth"
	"github.com/carlos0934/videotube/controllers"
	"github.com/carlos0934/videotube/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := BuildContainer()

	app.StartServer(":3000")

}

func BuildContainer() *controllers.AppServer {

	conn := db.MysqlConnector("root@/videotube")
	key := auth.GetECPrivateKey("key.pem", "public.pem")
	userController := controllers.NewUserController(conn, key)
	videoController := controllers.NewVideoController(conn)
	app := controllers.NewAppServer()
	app.AddRouter(userController, videoController)

	return app
}
