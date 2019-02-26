package stream

type Stream struct {
	Type     string      `json:"type,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Settings Settings    `json:"settings"`
}
