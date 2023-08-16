package main

import (
	"duolingo/server"
	"net/http"
)

func main() {
	server.AddStatusHandler()
	server.AddDiscordInteractionHandler()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
