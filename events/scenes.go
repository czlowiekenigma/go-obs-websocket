package events

import "github.com/czlowiekenigma/go-obs-websocket/scene/source"

type SwitchScenes struct {
	rawEvent
	SceneName string           `json:"scene-name"`
	Sources   []*source.Source `json:"sources"`
}

type ScenesChanged struct {
	rawEvent
}

type SceneCollectionChanged struct {
	rawEvent
}

type SceneCollectionListChanged struct {
	rawEvent
}
