package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"./handlers"

	"github.com/sandalwing/echo-logrusmiddleware"
)

// EchoServer is the API Server Object
var APIServer *echo.Echo

func createServer() {
	APIServer = echo.New()
	APIServer.HideBanner = true
	APIServer.Logger = logrusmiddleware.Logger{log.StandardLogger()}
	APIServer.Use(logrusmiddleware.Hook())
}

// Start the API
func Start(port string) {
	log.Fatal(APIServer.Start(":" + port))
}

// Initialize creates the API endpoints for the current service
func Initialize() {

	createServer()

	//
	// Register API Endpoinds here
	//

	// Example Tasks Endpoints
	APIServer.GET("/tasks", APIHandlers.GetTasks())
	APIServer.PUT("/tasks", APIHandlers.PutTask())
	APIServer.DELETE("/tasks/:id", APIHandlers.DeleteTask())
}
