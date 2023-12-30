package server

import (
	"duolingo/discord"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		panic(err)
	}
}

func AddStatusHandler() {
	http.HandleFunc("/status", StatusHandler)
}

func AddDiscordInteractionHandler() {
	http.HandleFunc("/interactions", func(w http.ResponseWriter, r *http.Request) {
		err := discord.VerifySignature(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("invalid request signature"))
			return
		}
		interaction := discord.DecodeInteraction(r)
		println("Received interaction of type", interaction.Type)
		switch interaction.Type {
		case 1:
			discord.AcknowledgeInteraction(w, interaction)
		case 2:
			discord.AnswerLeaderboard(w, interaction)
		}
	})
}
