package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"net/http"

	"github.com/gorilla/websocket"
)

var (
	socketLogStd = log.New(os.Stdout, "[socket] ", log.Ldate|log.Ltime)
	socketLogErr = log.New(os.Stderr, "ERROR [socket] ", log.Ldate|log.Ltime)

	clients = make(map[*websocket.Conn]bool)
	// broadcast = make(chan GameStatus)

	upgrader = websocket.Upgrader{
		CheckOrigin: func(request *http.Request) bool {
			return true
		},
	}
)

const (
	relPath = "/socket/"
	host    = ":80"
)

func handleConnections(response http.ResponseWriter, request *http.Request) {

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		socketLogErr.Printf("upgrade error: %v", err)
		return
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	fmt.Println(request.URL.Path)
	channel, err := strconv.Atoi(strings.TrimPrefix(request.URL.Path, relPath))
	if err != nil {
		socketLogErr.Printf("connection error: %v", err)
		return
	}
	if channel < 0 || channel > 2 {
		socketLogErr.Println("invalid channel:", channel)
		return
	}

	socketLogStd.Println(
		"New connection to channel",
		channel,
		"from",
		request.RemoteAddr,
	)

	clients[ws] = true

	for {
		//TODO: recieve input here and handle it
	}
}

func sendGameStatus() {
	// Send current game status out to every client that is currently connected
	for {
		for client := range clients {
			err := client.WriteJSON(gameStatus)
			if err != nil {
				socketLogErr.Printf("write error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}

	}
}
