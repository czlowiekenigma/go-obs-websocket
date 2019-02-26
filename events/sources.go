package events

type SourceOrderChanged struct {
	rawEvent
	SceneName string `json:"scene-name"`
}

type SceneItemAdded struct {
	rawEvent
	SceneName string `json:"scene-name"`
	ItemName  string `json:"item-name"`
}

type SceneItemRemoved struct {
	rawEvent
	SceneName string `json:"scene-name"`
	ItemName  string `json:"item-name"`
}

type SceneItemVisibilityChanged struct {
	rawEvent
	SceneName   string `json:"scene-name"`
	ItemName    string `json:"item-name"`
	ItemVisible bool   `json:"item-visible"`
}
