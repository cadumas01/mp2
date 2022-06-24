# mp2
 a simple chat room application that supports only private message


# Usage

- Send a message to another user with: ```send [DESTINATION_USERNAME] [MESSAGE]```

# Ideas / Notes
- Server contains map of all incoming connections and outgoing connections
    - All messagses go through server
    - Server has many incoming connections, and many outgoing connections 
        - We use two different maps so clients can have an outgoing channel and incoming channel
            - Use of gourtines becomes easier 
    - Each Client has one incoming connection and one outgoing connection (server)

- Server relaying message to destination:
    - First see if conn exits in map
    - If not, then search config file for port /address to dial
        - If not found, print error
    - If dial fails then destination is not connected (print error)

# Message Guidelines
- Message Content may not have new lines
- TO and FROM fields much match usernames of clients in config.json


- Servers and Clients have same IP (hostAddress) but unique ports
