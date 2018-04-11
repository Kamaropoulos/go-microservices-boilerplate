package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/sandalwing/echo-logrusmiddleware"
)

// EchoServer is the API Server Object
var EchoServer *echo.Echo

func createServer() {
	EchoServer = echo.New()

	EchoServer.Logger = logrusmiddleware.Logger{log.StandardLogger()}
	EchoServer.Use(logrusmiddleware.Hook())
}

// Start the API
func Start(port string) {
	log.Fatal(EchoServer.Start(":" + port))
}

// Initialize creates the API endpoints for the current service
func Initialize() {

	createServer()
	EchoServer.HideBanner = true

	//
	// Register API Endpoinds here
	//

}
