package client

import (
	"bufio"
	"bytes"
	"fmt"
	"mp2/configurations"
	"mp2/messages"
	"mp2/server"
	"net"
	"os"
	"strconv"
	"strings"
)

var config configurations.Config
var self *configurations.Client

func StartClient(username string, port int, hostAddress string, c configurations.Config) {
	config = c
	self = &configurations.Client{Username: username, Port: port, HostAddress: hostAddress}

	//0. listen and accept connection back from server
	ln := server.Listen(port)

	// 1. Connect to Server
	// 1.a. Send userame so server can add to map
	outConn := connectToServer()

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

func connectToServer() (outConn net.Conn) {
	serverPort, _ := configurations.QuerryConfig(config, "server", "")

	socket := self.HostAddress + ":" + strconv.Itoa(serverPort)

	//Connect to port
	outConn, err := net.Dial("tcp", socket)

	if err != nil {
		panic(err)
	}

	// write username to server
	outConn.Write([]byte(self.Username))

	fmt.Println("Client successfully connected to server")

	return outConn
}

// Handles sending messages
func handleCLI(outConn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//strip new line
	text = strings.Replace(text, "\n", "", -1)

	textArr := strings.Split(text, " ")

	if len(textArr) < 3 {
		sendUsage()
		return
	}

	// Build message
	destination := textArr[1]
	content := strings.Join(textArr[2:], " ")
	message := messages.NewMessage(destination, self.Username, content)

	// Send Message
	outConn.Write([]byte(message.String()))
	fmt.Println("Send message to server")
}

func sendUsage() {
	fmt.Println("Invalid command\nSend Message with: send [DESTINATION_USERNAME] [MESSAGE]")
}

// Handles incoming messages
func handleConnection(inConn net.Conn) {
	for {
		buf := make([]byte, server.BufSize)
		_, err := bufio.NewReader(inConn).Read(buf)

		// if err is empty, we have a message and can print it
		if err != nil {
			panic(err)
		}

		// Don't bother parsing messageString back to Message struct if we are simply printing it right away
		messageString := string(bytes.Trim(buf, "\x00")) //trims buf of empty bytes

		// Handles potential exit
		if messageString == server.Exit {
			fmt.Println("Exiting")
			os.Exit(0)
		}

		fmt.Println("Received a message:\n'" + messageString)
	}
}
