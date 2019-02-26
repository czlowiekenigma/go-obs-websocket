package filter

type Filter struct {
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Settings interface{} `json:"settings"`
}
