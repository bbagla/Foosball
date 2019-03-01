package main

import (
	serve "concurrency-17/src/server"
	"flag"
	"log"
	"net/http"
)

const (
	//SocketPath = "/socket/"
	resetPath  = "/reset/"
	host       = "localhost:8000"
)

func main() {
	hostPtr :=flag.String("host","localhost:80","host ptr")
	flag.Parse()
	http.HandleFunc("/socket/", serve.HandleConnections)
	http.HandleFunc(
		resetPath,
		func(response http.ResponseWriter, request *http.Request) {
			// resetGame()
			response.WriteHeader(http.StatusOK)
		},
	)

	serve.PrepareGame()

	log.Println("Listening on", *hostPtr)
	http.ListenAndServe(*hostPtr, nil)
}
