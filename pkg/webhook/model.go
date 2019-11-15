package webhook

import "github.com/LeviMatus/soundboard/types"

type Webhook struct {
	User  types.JiraUser `json:"user"`
	Issue Issue          `json:"issue"`
}

type Issue struct {
	Id     string `json:"id"`
	Link   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Points float32 `json:"customfield_10001"`
	Status Status  `json:"status"`
}

type Status struct {
	Status string `json:"name"`
}
