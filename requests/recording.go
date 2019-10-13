package requests

type StartStopRecording struct {
	RequestBase
}

type StartRecording struct {
	RequestBase
}

type StopRecording struct {
	RequestBase
}

type PauseRecording struct {
	RequestBase
}

type ResumeRecording struct {
	RequestBase
}

type SetRecordingFolder struct {
	RequestBase
	RecFolder string `json:"rec-folder"`
}