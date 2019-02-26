package responses

import "github.com/defabricated/go-obs-websocket/stream"

type GetStreamingStatus struct {
	*ResponseBase
	Streaming      bool   `json:"streaming"`
	Recording      bool   `json:"recording"`
	StreamTimecode string `json:"stream-timecode,omitempty"`
	RecordTimecode string `json:"rec-timecode,omitempty"`
	PreviewOnly    bool   `json:"preview_only"`
}

type GetStreamSettings struct {
	*ResponseBase
	Type     stream.Type     `json:"type"`
	Settings stream.Settings `json:"settings"`
}
