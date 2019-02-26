package responses

import "github.com/defabricated/go-obs-websocket/transition"

type GetTransitionList struct {
	*ResponseBase
	CurrentTransition string                   `json:"current-transition"`
	Transitions       []*transition.Transition `json:"transitions"`
}

type GetCurrentTransition struct {
	*ResponseBase
	transition.Transition
}

type GetTransitionDuration struct {
	*ResponseBase
	Duration int `json:"transition-duration"` //TODO Change to time.Time
}
