package configurations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Server  Server   `json:"server"`
	Clients []Client `json:"clients"`
}

type Server struct {
	Port int `json:"port"`
}

type Client struct {
	Username    string `json:"username"`
	Port        int    `json:"port"`
	HostAddress string `json:"hostAddress"`
}

func GetConfig() Config {
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

func QuerryConfig(config Config, processType string, clientUsername string) (port int, hostAddress string) {
	if &config == nil {
		return 0, ""
	}
	if processType == "server" {
		port := config.Server.Port
		return port, ""
	}
	if processType == "client" {
		port, hostAddress := querryClients(config.Clients, clientUsername)
		return port, hostAddress
	}
	return 0, ""
}

func querryClients(clients []Client, username string) (port int, hostAddress string) {
	for _, client := range clients {
		if client.Username == username {
			return client.Port, client.HostAddress
		}
	}
	return 0, ""
}
