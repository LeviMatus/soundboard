package main

import (
	"encoding/json"
	"fmt"
	"github.com/LeviMatus/soundboard/pkg/sounds"
	"github.com/LeviMatus/soundboard/pkg/webhook"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var samples = make(chan string, 10)
var users []sounds.User

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("got here")

	var b webhook.Webhook
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range users {
		if v.IsUser(b.User) {
			samples <- v.ClipName
		}
	}

	fmt.Println("got webhook payload:")
	fmt.Println(b.Issue.Fields.Points, b.User.Key, b.Issue.Fields.Status.Status)
}

func main() {

	var v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("sounds.yaml")
	v.AddConfigPath("/home/levi/GolandProjects/soundboard")
	v.ReadInConfig()

	err := v.Unmarshal(&users)

	if err != nil {
		fmt.Println(err)
	}

	go sounds.Consume(samples)
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
