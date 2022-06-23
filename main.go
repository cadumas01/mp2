package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var config Config

type Config struct {
	server  Server   `json:"server`
	clients []Client `json:"clients"`
}

type Server struct {
	port int `json:"port"`
}

type Client struct {
	username    string `json:"username"`
	port        int    `json:"port"`
	hostAddress string `json:"hostAddress"`
}

func usage() {
	fmt.Println("USAGE:\n\tTo start server (do this first): go run main.go server\n\tTo start a node: go run main.go [CLIENT_USERNAME]\n\t**CLIENT_USERNAME must be with client in config.json")
}

func getConfig() Config {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	// Must unmarshall the json object
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c Config
	json.Unmarshal(byteValue, &c)

	jsonFile.Close()

	return c
}

func querryConfig(config Config, processType string, clientUsername string) (port int, hostAddress string) {
	if &config == nil {
		return 0, ""
	}
	if processType == "server" {
		port := config.server.port
		return port, ""
	}
	if processType == "client" {
		port, hostAddress := querryClients(config.clients, clientUsername)
		return port, hostAddress
	}
	return 0, ""
}

func querryClients(clients []Client, username string) (port int, hostAddress string) {
	for _, client := range clients {
		if client.username == username {
			return client.port, client.hostAddress
		}
	}
	return 0, ""
}

func main() {
	config = getConfig()

	// Starting Node / CLI handling
	args := os.Args
	if len(args) == 2 && args[1] == "Server" {
		// start server
	} else if len(args) == 3 && args[1] == "Client" {
		username := args[2]
		port, hostAddress := querryConfig(config, "client", username)
		if port == 0 {
			fmt.Println("Client username not found in config")
			usage()
			os.Exit(1)
		}
		// START CLIENT

	}

	usage()
	os.Exit(1)

}
