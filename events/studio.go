package events

import "github.com/defabricated/go-obs-websocket/scene/source"

type PreviewSceneChanged struct {
	rawEvent
	SceneName string           `json:"scene-name"`
	Sources   []*source.Source `json:"sources"`
}

type StudioModeSwitched struct {
	rawEvent
	NewState bool `json:"new-state"`
}
