package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/LeviMatus/soundboard/pkg/sounds"
	"github.com/LeviMatus/soundboard/types"
	"log"
	"net/http"
)

var samples = make(chan string, 10)
var users []types.SoundMap

func Listen(u []types.SoundMap) {
	users = u
	go sounds.Consume(samples)
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var b Webhook
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range users {
		if v.IsJiraUser(b.User) {
			samples <- v.ClipName
		}
	}

	fmt.Println("got webhook payload:")
	fmt.Println(b.Issue.Fields.Points, b.User.Key, b.Issue.Fields.Status.Status)
}
