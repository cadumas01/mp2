# mp2
 a simple chat room application that supports only private message


# Usage

- Send a message to another user with: ```send [DESTINATION_USERNAME] [MESSAGE]```

# Model
- One server, many clients
- To pass a message from client A to client B:
    - Client A sends message to server
    - Server sends message to Client B

- Each client has two connections with the server: incoming and outgoing
    - Each are on separate goroutines
    - Outgoing channel interpret message sending commands, sending messages to server
    - Incoming channel interprets incoming text and displays messages to stdout

- Server has a goroutine for each incoming and outgoing channel per client connection
    - Incoming channels interpret messages (looking at "To" field) to relay message to final destination using outgoing channels
        - Server manages incoming and outgoing connections with maps: mapping client username to corresponding connection instance
    - Goroutines allow server to listen for inputs from (and relay messagest to) all clients concurrently
    - Server checks to make sure message destination exits and is connected, and handles error otherwise
- Server has separate gourtine listening to it's stdin for exit command ```EXIT``` to exit all processes


# Message Guidelines
- Message content may not have new lines
- TO and FROM fields much match usernames of clients in config.json

# Config Guidelines
- Servers and clients have same IP (hostAddress) but must have unique ports
- Ports may not be 0

# Resources

[Assignment instructions](https://docs.google.com/document/d/1gTO2W30h6_OWS-DFlSgiaMMHuyShgG3OBxVbcWnb1Ug/edit) 

Modeled after: [mp1](https://github.com/cadumas01/mp1)

