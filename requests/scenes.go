package requests

import "github.com/czlowiekenigma/go-obs-websocket/scene/item"

type SetCurrentScene struct {
	RequestBase
	SceneName string `json:"scene-name"`
}

type GetCurrentScene struct {
	RequestBase
}

type ReorderSceneItems struct {
	RequestBase
	Scene string       `json:"scene,omitempty"`
	Items []item.Item `json:"items"`
}
