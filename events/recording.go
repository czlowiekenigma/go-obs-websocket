package events

type RecordingStarting struct {
	rawEvent
}

type RecordingStarted struct {
	rawEvent
}

type RecordingStopping struct {
	rawEvent
}

type RecordingStopped struct {
	rawEvent
}

type RecordingPaused struct {
	rawEvent
}

type RecordingResumed struct {
	rawEvent
}