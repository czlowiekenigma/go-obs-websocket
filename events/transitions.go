package events

import "github.com/defabricated/go-obs-websocket/transition"

type SwitchTransition struct {
	rawEvent
	Name string `json:"transition-name"`
}

type TransitionListChanged struct {
	rawEvent
}

type TransitionDurationChanged struct {
	rawEvent
	NewDuration int `json:"new-duration"` //TODO Change to time.Duration
}

type TransitionBegin struct {
	rawEvent
	transition.Transition
	FromScene string `json:"from-scene"`
	ToScene string `json:"to-scene"`
}
