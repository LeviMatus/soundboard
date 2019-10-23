package sounds

import "github.com/LeviMatus/soundboard/pkg/webhook"

type User struct {
	Name     string
	ClipName string
}

func (u User) IsUser(ju webhook.User) bool {
	return u.Name == ju.Name
}
