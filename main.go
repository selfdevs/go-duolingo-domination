package main

import (
	"duolingo/discord"
	"net/http"
)

func main() {
	http.HandleFunc("/interactions", func(w http.ResponseWriter, r *http.Request) {
		err := discord.VerifySignature(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("invalid request signature"))
			return
		}
		interaction := discord.DecodeInteraction(r)
		println("Got interaction", interaction.Type)
		if interaction.Type == 1 {
			discord.AcknowledgeInteraction(w, interaction)
			return
		}
		if interaction.Type == 2 {
			discord.AnswerHealthCheck(w, interaction)
			return
		}
		return
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err)
		return
	}
}
