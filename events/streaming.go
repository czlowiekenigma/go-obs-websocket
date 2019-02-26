package events

type StreamStarting struct {
	rawEvent
	PreviewOnly bool `json:"preview-only"`
}

type StreamStarted struct {
	rawEvent
}

type StreamStopping struct {
	rawEvent
	PreviewOnly bool `json:"preview-only"`
}

type StreamStopped struct {
	rawEvent
}

type StreamStatus struct {
	rawEvent
	Streaming               bool    `json:"streaming"`
	Recording               bool    `json:"recording"`
	PreviewOnly             bool    `json:"preview-only"`
	BytesPerSec             int     `json:"bytes-per-sec"`
	KBitsPerSec             int     `json:"kbits-per-sec"`
	DroppedFramesPercentage float64 `json:"strain"`
	TotalStreamTime         int     `json:"total-stream-time"` //TODO Change to time.Duration
	NumTotalFrames          int     `json:"num-total-frames"`
	NumDroppedFrames        int     `json:"num-dropped-frames"`
	FPS                     float64 `json:"fps"`
}
