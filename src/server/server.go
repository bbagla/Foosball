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
	keyboardInput KeyboardInput

	upgrader = websocket.Upgrader{
		CheckOrigin: func(request *http.Request) bool {
			return true
		},
	}
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
	channel, err := strconv.Atoi(strings.TrimPrefix(request.URL.Path, socketPath))
	if err != nil {
		socketLogErr.Printf("connection error: %v", err)
		return
	}
	if channel < 0 || channel > 1 {
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
		err := ws.ReadJSON(&keyboardInput)
		if err != nil {
			socketLogErr.Printf("read error: %v", err)
			ws.Close()
			delete(clients, ws)
			return
		}
		if keyboardInput.KeyPressed == 3 {
			startGame()
			continue
		}
		//else reset game

		if channel != -1 {
			var stick1 = gameStatus.Teams[channel].GoalKeeper[0:1]
			var stick2 = gameStatus.Teams[channel].Defence[0:2]
			var stick3 = gameStatus.Teams[channel].Mid[0:5]
			var stick4 = gameStatus.Teams[channel].Attack[0:3]
			if keyboardInput.SelectStick == 1 {
				gameStatus.Teams[channel].LastStick = stick1
			} else if keyboardInput.SelectStick == 2 {
				gameStatus.Teams[channel].LastStick = stick2
			} else if keyboardInput.SelectStick == 3 {
				gameStatus.Teams[channel].LastStick = stick3
			} else if keyboardInput.SelectStick == 4 {
				gameStatus.Teams[channel].LastStick = stick4
			}
		}
	}
}

func sendGameStatus() {
	// Send current game status out to every client that is currently connected
	for client := range clients {
		err := client.WriteJSON(gameStatus)
		if err != nil {
			socketLogErr.Printf("write error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
