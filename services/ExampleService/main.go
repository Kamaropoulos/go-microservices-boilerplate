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

	// Initialize Logger
	initLogger()

	// Log startup messages
	log.Info("Service Name: " + ServiceName)
	log.Info("Starting API...")

	// Get port for the service to run on
	var servicePort string

	// Read service port from the first command line argument
	if len(os.Args) > 1 {

		// TODO: Check if the port number passed is a valid port number.
		// 		 Echo will do that anyways but it'd be a good thing to do.
		servicePort = os.Args[1]

		if servicePort == "" {
			log.Fatal("Service port number was not set.")
		}

		log.Debug("Service Port set to " + servicePort)

	} else {
		// No arguments passed, service can't run without a port
		log.Fatal("Service port number was not set.")
	}

	// Get address and port of the Service Registry (if it exists)

	// Variable to store the Registry Service state
	usingRegistry := false
	_ = usingRegistry

	var registryAddress string
	var registryPort string

	// If there are more than 3 arguments this means that an address and port were passed
	if len(os.Args) > 3 {

		// Store arguments
		// TODO: Verify if the arguments are valid and the service reachable
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
		// There are no more than 3 arguments and the service registry is
		// considered to be running on localhost listening at the passed port
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
			// If there are no more than two arguments and since we passed the servicePort checks
			// there should only be the servicePort argument.
			// As a result, Service Registry won't be used
			usingRegistry = false
			log.WithFields(log.Fields{
				"usingRegistry": usingRegistry,
			}).Warning("Service Registry address was not set.")
		}
	}

	// TODO: Connect to DB

	// TODO: Connect to Service Registry if connection information were passed

	//Initialize API
	api.Initialize()

	// Start API
	//api.Start(port)
	api.Start(servicePort)

	// TODO: Init RPC
	// TODO: Register Service

}
