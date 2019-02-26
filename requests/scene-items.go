package requests

import (
	"github.com/defabricated/go-obs-websocket/scene"
	"github.com/defabricated/go-obs-websocket/transforms"
	"github.com/guregu/null"
)

type GetSceneItemProperties struct {
	RequestBase
	SceneName string `json:"scene-name,omitempty"`
	Item      string `json:"item"`
}

type SetSceneItemProperties struct {
	RequestBase
	SceneName string               `json:"scene-name,omitempty"`
	Item      string               `json:"item"`
	Position  *transforms.Position `json:"position,omitempty"`
	Rotation  float64              `json:"rotation,omitempty"`
	Scale     *transforms.Scale    `json:"scale,omitempty"`
	Crop      *transforms.Crop     `json:"crop,omitempty"`
	Visible   null.Bool            `json:"visible,omitempty"`
	Bounds    *transforms.Bounds   `json:"bounds,omitempty"`
}

type ResetSceneItem struct {
	RequestBase
	SceneName string `json:"scene-name"`
	Item      string `json:"item"`
}

type DeleteSceneItem struct {
	RequestBase
	Scene string     `json:"scene"`
	Item  scene.Item `json:"item"`
}

type DuplicateSceneItem struct {
	RequestBase
	FromScene string     `json:"from-scene,omitempty"`
	ToScene   string     `json:"to-scene,omitempty"`
	Item      scene.Item `json:"item"`
}
