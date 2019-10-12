package routes

import (
	"github.com/daitasu/youtube-manager-go/web/api" // 1
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.start(":8080"))
}