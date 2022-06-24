package configurations

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Server  Server   `json:"server"`
	Clients []Client `json:"clients"`
}

type Server struct {
	Port int `json:"portServer"`
}

type Client struct {
	Username    string `json:"username"`
	Port        int    `json:"port"`
	HostAddress string `json:"hostAddress"`
}

func GetConfig() Config {
	jsonFile, err := os.Open("configurations/config.json")

	if err != nil {
		panic(err)
	}

	// Must unmarshall the json object
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

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

func Exists(config Config, processType string, clientUsername string) bool {
	port, _ := QuerryConfig(config, processType, clientUsername)
	return port != 0
}
