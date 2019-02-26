package requests

type SetCurrentTransition struct {
	RequestBase
	Name string `json:"transition-name"`
}

type SetTransitionDuration struct {
	RequestBase
	Duration int `json:"duration"` //TODO Change to time.Time
}