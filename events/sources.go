package events

import (
	"github.com/czlowiekenigma/go-obs-websocket/scene/item"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/filter"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/mixer"
)

type SourceCreated struct {
	rawEvent
	Name     string      `json:"sourceName"`
	Type     string      `json:"sourceType"`
	Kind     string      `json:"sourceKind"`
	Settings interface{} `json:"sourceSettings"`
}

type SourceDestroyed struct {
	rawEvent
	Name string `json:"sourceName"`
	Type string `json:"sourceType"`
	Kind string `json:"sourceKind"`
}

type SourceVolumeChanged struct {
	rawEvent
	Name   string  `json:"sourceName"`
	Volume float64 `json:"volume"`
}

type SourceMuteStateChanged struct {
	rawEvent
	Name  string `json:"sourceName"`
	Muted bool   `json:"muted"`
}

type SourceAudioSyncOffsetChanged struct {
	rawEvent
	Name       string `json:"sourceName"`
	SyncOffset int    `json:"syncOffset"` //TODO Convert to time.Duration
}

type SourceAudioMixersChanged struct {
	rawEvent
	Name           string         `json:"sourceName"`
	Mixers         []*mixer.Mixer `json:"mixers"`
	HexMixersValue string         `json:"hexMixersValue"`
}

type SourceRenamed struct {
	rawEvent
	PreviousName string `json:"previousName"`
	NewName      string `json:"newName"`
}

type SourceFilterAdded struct {
	rawEvent
	Name           string      `json:"sourceName"`
	FilterName     string      `json:"filterName"`
	FilterType     string      `json:"filterType"`
	FilterSettings interface{} `json:"filterSettings"`
}

type SourceFilterRemoved struct {
	rawEvent
	Name       string `json:"sourceName"`
	FilterName string `json:"filterName"`
	FilterType string `json:"filterType"`
}

type SourceFiltersReordered struct {
	rawEvent
	Name    string           `json:"sourceName"`
	Filters []*filter.Filter `json:"filters"`
}

type SourceOrderChanged struct {
	rawEvent
	SceneName  string `json:"scene-name"`
	SceneItems []*struct {
		Name string `json:"source-name"`
		ID   int    `json:"item-id"`
	} `json:"scene-items"`
}

type SceneItemAdded struct {
	rawEvent
	SceneName string `json:"scene-name"`
	Name      string `json:"item-name"`
	ID        int    `json:"item-id"`
}

type SceneItemRemoved struct {
	rawEvent
	SceneName string `json:"scene-name"`
	Name      string `json:"item-name"`
	ID        int    `json:"item-id"`
}

type SceneItemVisibilityChanged struct {
	rawEvent
	SceneName string `json:"scene-name"`
	Name      string `json:"item-name"`
	ID        int    `json:"item-id"`
	Visible   bool   `json:"item-visible"`
}

type SceneItemTransformChanged struct {
	rawEvent
	SceneName string                   `json:"scene-name"`
	Name      string                   `json:"item-name"`
	ID        int                      `json:"item-id"`
	Transform *item.SceneItemTransform `json:"transform"`
}

type SceneItemSelected struct {
	rawEvent
	SceneName string `json:"scene-name"`
	Name      string `json:"item-name"`
	ID        int    `json:"item-id"`
}

type SceneItemDeselected struct {
	rawEvent
	SceneName string `json:"scene-name"`
	Name      string `json:"scene-name"`
	ID        int    `json:"item-id"`
}
