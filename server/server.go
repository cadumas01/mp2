package server

import (
	"bufio"
	"bytes"
	"fmt"
	"mp2/configurations"
	"mp2/messages"
	"net"
	"os"
	"strconv"
	"strings"
)

const BufSize int = 1024

var config configurations.Config

const Exit string = "EXIT"

func StartServer(port int, c configurations.Config) {

	config = c

	ln := Listen(port)

	// maps node username to net.Conn (that client to server)
	inConns := make(map[string]net.Conn)
	outConns := make(map[string]net.Conn)

	// handles input from server CLI
	go acceptClients(inConns, outConns, ln)

	handleServer(outConns)
}

func Listen(port int) (ln net.Listener) {
	address := ":" + strconv.Itoa(port)
	ln, err := net.Listen("tcp", address)

	if err != nil {
		panic("Error listening")
	}

	fmt.Println("server started for " + ln.Addr().String())

	return ln
}

//NEED TO ADJUST?
// Waits for client to connect and recieves message
func acceptClients(inConns map[string]net.Conn, outConns map[string]net.Conn, ln net.Listener) {
	// loop to allow function to accept all clients
	for {
		conn, err := ln.Accept()

		if err != nil {
			panic("error accepting")
		}

		// Get accepted id to add to list
		buf := make([]byte, BufSize)
		_, err = bufio.NewReader(conn).Read(buf)

		// if err is empty, we have a message and can print it
		if err != nil {
			panic(err)
		}

		port := conn.RemoteAddr().String()            // temp?
		fmt.Println("\naccepted port = " + port)      // temp?
		acceptedUn := string(bytes.Trim(buf, "\x00")) //trims buf of empty bytes

		// Un = client username
		inConns[acceptedUn] = conn
		fmt.Println("Just accepted username = " + acceptedUn)

		// Once incoming conn has been added, immediately dial back to client to establish outgoing client
		// Or maybe dial only when needed
		connectTo(acceptedUn, outConns) // maybe add goroutine?

		// Handles incoming messages (and redirects them with outConns to destination client)
		go handleConnection(conn, outConns)
	}
}

func connectTo(username string, outConns map[string]net.Conn) {
	port, hostAddress := configurations.QuerryConfig(config, "client", username)

	socket := hostAddress + ":" + strconv.Itoa(port)

	//Connect to port
	conn, err := net.Dial("tcp", socket)

	if err != nil {
		panic(err)
	}

	// add to map - NOT DONE
	outConns[username] = conn

	fmt.Println("Server successfully connected back to  " + username)
}

// Handles incoming messages for the node
func handleConnection(conn net.Conn, outConns map[string]net.Conn) {
	// loop to allow for many connection handling
	for {
		buf := make([]byte, BufSize)
		_, err := bufio.NewReader(conn).Read(buf)

		if err != nil {
			panic(err)
		}

		message := messages.ParseMessage(string(bytes.Trim(buf, "\x00"))) //trims buf of empty bytes and parse to Message
		relayMessage(message, outConns)
	}
}

// Relays message to final destination
func relayMessage(message *messages.Message, outConns map[string]net.Conn) {
	// if message.To is in outConns...
	if conn, ok := outConns[message.To]; ok {
		fmt.Println("\nRelaying a message from '" + message.From + "' to '" + message.To + "' \n")
		_, err := conn.Write([]byte(message.String()))

		if err != nil {
			panic("error writing to destination")
		}
	} else {
		fmt.Println("destination not found. Sending notification to sender")

		fromConn := outConns[message.From]

		_, err := fromConn.Write([]byte("Cannot send message. '" + message.To + "' is not connected."))
		if err != nil {
			panic("error writing back to from-client")
		}
	}
}

// Handles Server CLI (deals with exit commands)
func handleServer(outConns map[string]net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1) // trim newline
		if input == Exit {
			exitAll(outConns)
		}
	}
}

func exitAll(outConns map[string]net.Conn) {

	fmt.Println("Exiting clients...")
	// send exit command to all clients
	for _, conn := range outConns {
		_, err := conn.Write([]byte(Exit))
		if err != nil {
			panic("error writing exit command")
		}
	}

	// exit server
	fmt.Println("Exiting server...")
	os.Exit(0)
}
