package duolingo

import (
	"duolingo/discord"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
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
}
