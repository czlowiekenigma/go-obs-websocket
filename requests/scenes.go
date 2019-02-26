package requests

import "github.com/defabricated/go-obs-websocket/scene"

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
	Items []scene.Item `json:"items"`
}
