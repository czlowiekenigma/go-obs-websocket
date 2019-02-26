package responses

import "github.com/defabricated/go-obs-websocket/scene/source"

type GetStudioModeStatus struct {
	*ResponseBase
	StudioMode bool `json:"studio-mode"`
}

type GetPreviewScene struct {
	*ResponseBase
	Name    string           `json:"name"`
	Sources []*source.Source `json:"sources"`
}
