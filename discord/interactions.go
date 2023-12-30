package discord

import (
	"bytes"
	"duolingo/canvas"
	"duolingo/duolingo"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type InteractionData struct {
	Content string `json:"content"`
}

type Interaction struct {
	Type      int              `json:"type"`
	ChannelId string           `json:"channel_id"`
	Data      *InteractionData `json:"data,omitempty"`
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

func generateAndAsyncSendImage(interaction Interaction) {
	users := duolingo.FetchUsers()
	canvas.DrawLeaderboard(users)
	println("Preparing to send the image")
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	file, openErr := os.Open("leaderboard.png")
	if openErr != nil {
		println("Fail to open file")
		os.Exit(1)
	}
	defer file.Close()

	field, err := writer.CreateFormFile("file", "leaderboard.png")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(field, file); err != nil {
		log.Fatal(err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "https://discord.com/api/v10/channels/"+interaction.ChannelId+"/messages", &buf)
	if err != nil {
		println("Fail to create request")
		os.Exit(1)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bot "+os.Getenv("DISCORD_BOT_TOKEN"))
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		println("Fail to do request")
		os.Exit(1)
	}
	println("Sent the image")
}

func AnswerLeaderboard(w http.ResponseWriter, interaction Interaction) {
	go generateAndAsyncSendImage(interaction)
	w.Header().Set("Content-Type", "application/json")
	interaction = Interaction{
		Type: 4,
		Data: &InteractionData{Content: "Sending it right away! (It might take a minute or two for me to process, I'm just an owl after all)"},
	}
	err := json.NewEncoder(w).Encode(interaction)
	if err != nil {
		println("Fail to answer health check")
		os.Exit(1)
	}
	println("Answered the interaction")
}
