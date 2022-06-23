package client

import (
	"fmt"
	"mp2/configurations"
	"mp2/server"
	"net"
	"strconv"
)

var config configurations.Config

func StartClient(username string, port int, hostAddress string, c configurations.Config) {
	config = c

	//0. listen and accept connection back from server
	ln := server.Listen(port)

	// 1. Connect to Server
	// 1.a. Send userame so server can add to map
	outConn := connectToServer(username, hostAddress)

	// Maybe goroutine and
	// Handle Server's return dial
	inConn, err := ln.Accept()

	if err != nil {
		panic("error accepting")
	}

	fmt.Println("Server successfully reconnected")

	// 2. (goroutine) Be available to send messages
	go handleCLI(outConn)

	// 3. (goroutine) Be available to receive messages
	go handleConnection(inConn)
}

func connectToServer(username string, hostAddress string) (outConn net.Conn) {
	serverPort, _ := configurations.QuerryConfig(config, "server", "")

	socket := hostAddress + ":" + strconv.Itoa(serverPort)

	//Connect to port
	outConn, err := net.Dial("tcp", socket)

	if err != nil {
		panic(err)
	}

	// write username to server
	outConn.Write([]byte(username))

	fmt.Println("Client successfully connected to server")

	return outConn
}

// ADD TO THESE
func handleCLI(outConn net.Conn) {

}

func handleConnection(inConn net.Conn) {
	// Parse incoming messages
	// If exit code, exit
	// Else, print message

}
