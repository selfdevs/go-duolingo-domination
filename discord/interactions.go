package discord

import (
	"encoding/json"
	"net/http"
	"os"
)

type InteractionData struct {
	Content string `json:"content"`
}

type Interaction struct {
	Type int             `json:"type"`
	Data InteractionData `json:"data"`
}

func DecodeInteraction(r *http.Request) Interaction {
	var interaction Interaction
	err := json.NewDecoder(r.Body).Decode(&interaction)
	if err != nil {
		println("Fail to decode interaction")
		os.Exit(1)
	}
	return interaction
}

func AcknowledgeInteraction(w http.ResponseWriter, interaction Interaction) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(interaction)
	if err != nil {
		println("Fail to acknowledge interaction")
		os.Exit(1)
	}
	println("ACK successful")
}

func AnswerHealthCheck(w http.ResponseWriter, interaction Interaction) {
	w.Header().Set("Content-Type", "application/json")
	interaction = Interaction{
		Type: 4,
		Data: InteractionData{Content: "I'm alive!"},
	}
	err := json.NewEncoder(w).Encode(interaction)
	if err != nil {
		println("Fail to answer health check")
		os.Exit(1)
	}
	println("Answered health check")
}
