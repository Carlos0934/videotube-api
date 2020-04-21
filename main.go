package main

import (
	"github.com/carlos0934/videotube/controllers"
)

func main() {
	app := controllers.NewAppServer()

	app.StartServer(":3000")
}
