package requests

import "github.com/defabricated/go-obs-websocket/transition"

type SetPreviewScene struct {
	RequestBase
	SceneName string `json:"scene-name"`
}

type TransitionToProgram struct {
	RequestBase
	WithTransition *transition.Transition `json:"with-transition"`
}

type EnableStudioMode struct {
	RequestBase
}

type DisableStudioMode struct {
	RequestBase
}

type ToggleStudioMode struct {
	RequestBase
}