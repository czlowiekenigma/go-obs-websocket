package transition

type Transition struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"` //TODO Change to time.Duration
}
