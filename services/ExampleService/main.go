package main

import (
	"os"

	"./api"

	log "github.com/Sirupsen/logrus"
)

// ServiceName The name of the service
var ServiceName = "ExampleService"

func initLogger() {
	// TODO: Move Log related stuff to another file
	log.SetLevel(log.DebugLevel)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}

func main() {

	initLogger()

	log.Info("Service Name: " + ServiceName)
	log.Info("Starting API...")

	// Get service port and registry address
	var port string

	if len(os.Args) > 1 {

		port = os.Args[1]

		if port == "" {
			log.Fatal("Service port number was not set.")
		}

		log.Debug("Service Port set to " + port)

	} else {
		log.Fatal("Service port number was not set.")
	}

	usingRegistry := false
	_ = usingRegistry

	var registryAddress string
	var registryPort string

	if len(os.Args) > 3 {

		registryAddress = os.Args[2]
		registryPort = os.Args[3]

		if (registryAddress == "") || (registryPort == "") {
			usingRegistry = false
			log.WithFields(log.Fields{
				"usingRegistry": usingRegistry,
			}).Warning("Service Registry address was not set.")
		}

		usingRegistry = true

		log.WithFields(log.Fields{
			"usingRegistry": usingRegistry,
		}).Debug("Service Registry address set to " + registryAddress)

		log.Debug("Service Registry Port set to " + registryPort)

	} else {
		if len(os.Args) > 2 {
			registryPort = os.Args[2]

			if registryPort == "" {
				usingRegistry = false
				log.Warning("Service Registry address was not set.")
			}

			registryAddress = "localhost"
			usingRegistry = true

			log.Debug("Service Registry Address set to localhost")

			log.WithFields(log.Fields{
				"usingRegistry": usingRegistry,
			}).Debug("Service Registry Port set to " + registryPort)
		} else {
			usingRegistry = false
			log.WithFields(log.Fields{
				"usingRegistry": usingRegistry,
			}).Warning("Service Registry address was not set.")
		}
	}

	// registry_port := os.Args[2]

	// TODO: Connect to DB

	//Initialize API
	api.Initialize()

	// Start API
	//api.Start(port)
	api.Start(port)

	// Init RPC
	// Register Service

}
