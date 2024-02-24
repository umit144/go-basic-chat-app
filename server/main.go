package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) { response.WriteHeader(http.StatusOK) })
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Soket server started on : 8080")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		panic("Error starting server: " + error.Error())
	}
}

func handleConnections(response http.ResponseWriter, request *http.Request) {
	connection, error := upgrader.Upgrade(response, request, nil)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer connection.Close()

	clients[connection] = true

	for {
		var message Message
		error := connection.ReadJSON(&message)
		if error != nil {
			fmt.Println(error)
			delete(clients, connection)
			return
		}

		broadcast <- message
	}
}

func handleMessages() {
	for {
		message := <-broadcast

		for client := range clients {
			error := client.WriteJSON(message)
			if error != nil {
				fmt.Println(error)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
