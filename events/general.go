package events

type Heartbeat struct {
	rawEvent
	Pulse             bool   `json:"pulse"`
	CurrentProfile    string `json:"current-profile,omitempty"`
	CurrentScene      string `json:"current-scene,omitempty"`
	Streaming         bool   `json:"streaming,omitempty,omitempty"`
	TotalStreamTime   int    `json:"total-stream-time,omitempty"` //TODO Change to time.Duration
	TotalStreamBytes  int    `json:"total-stream-bytes,omitempty"`
	TotalStreamFrames int    `json:"total-stream-frames,omitempty"`
	Recording         bool   `json:"recording,omitempty"`
	TotalRecordTime   int    `json:"total-record-time,omitempty"` //TODO Change to time.Duration
	TotalRecordBytes  int    `json:"total-record-bytes,omitempty"`
	TotalRecordFrames int    `json:"total-record-frames,omitempty"`
}

type BroadcastCustomMessage struct {
	rawEvent
	Realm string      `json:"realm"`
	Data  interface{} `json:"data"`
}
