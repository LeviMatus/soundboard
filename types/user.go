package types

type JiraUser struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	DispName string `json:"displayName"`
}

type SoundMap struct {
	Name     string `mapstructure:"name", yaml:"name"`
	ClipName string `mapstructure:"clip", yaml:"clip"`
}

type User struct {
	JiraUser   JiraUser
	ConfigUser SoundMap
}

func (u SoundMap) IsJiraUser(ju JiraUser) bool {
	return u.Name == ju.Name
}
