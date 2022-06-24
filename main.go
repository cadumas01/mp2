package main

import (
	"fmt"
	"mp2/client"
	"mp2/configurations"
	"mp2/server"
	"os"
)

var config configurations.Config

func usage() {
	fmt.Println("USAGE:\n\tTo start server (do this first): go run main.go server\n\tTo start a node: go run main.go client [CLIENT_USERNAME]\n\t**CLIENT_USERNAME must be with client in config.json")
}

func main() {
	config = configurations.GetConfig()

	// Starting Node / CLI handling
	args := os.Args
	if len(args) == 2 && args[1] == "server" {
		port, _ := configurations.QuerryConfig(config, "server", "")
		server.StartServer(port, config)

	} else if len(args) == 3 && args[1] == "client" {
		username := args[2]
		port, hostAddress := configurations.QuerryConfig(config, "client", username)
		if port == 0 {
			fmt.Println("Client username not found in config")
			usage()
			os.Exit(1)
		}
		client.StartClient(username, port, hostAddress, config)

	} else {
		usage()
		os.Exit(1)
	}
}
