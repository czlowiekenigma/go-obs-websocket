package responses

import "github.com/czlowiekenigma/go-obs-websocket/scene"

type GetCurrentScene struct {
	*ResponseBase
	scene.Scene
}

type GetSceneList struct {
	*ResponseBase
	CurrentScene string         `json:"current-scene"`
	Scenes       []*scene.Scene `json:"scenes"`
}
