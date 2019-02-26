package transforms

type Bounds struct {
	Type      string  `json:"type"`
	Alignment int     `json:"alignment"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
}
