package source

type SourceType string

const (
	Input SourceType = "input"
	Filter SourceType = "filter"
	Transition SourceType = "transition"
	Scene SourceType = "scene"
	Unknown SourceType = "unknown"
)
