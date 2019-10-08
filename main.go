package main

import (
	"encoding/json"
	"fmt"
	"github.com/LeviMatus/soundboard/pkg/sounds"
	"github.com/LeviMatus/soundboard/pkg/webhook"
	"log"
	"net/http"
)

var samples = make(chan string, 10)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("got here")
	samples <- "ending.mp3"
	var b webhook.Webhook
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("got webhook payload:")
	fmt.Println(b.Issue.Fields.Points, b.User.Key, b.Issue.Fields.Status.Status)
}

func main() {
	go sounds.Consume(samples)
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
