package main

import (
	"github.com/carlos0934/videotube/controllers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := controllers.NewAppServer()

	app.StartServer(":3000")
}
