package webhook

type Webhook struct {
	User  User  `json:"user"`
	Issue Issue `json:"issue"`
}

type User struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	DispName string `json:"displayName"`
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
