package events

type ReplayStarting struct {
	rawEvent
}

type ReplayStarted struct {
	rawEvent
}

type ReplayStopping struct {
	rawEvent
}

type ReplayStopped struct {
	rawEvent
}
