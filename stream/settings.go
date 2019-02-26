package stream

type Settings struct {
	Server   string `json:"server,omitempty"`
	Key      string `json:"key,omitempty"`
	UseAuth  string `json:"use-auth,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
