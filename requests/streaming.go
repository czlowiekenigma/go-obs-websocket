package requests

import "github.com/czlowiekenigma/go-obs-websocket/stream"

type StartStopStreaming struct {
	RequestBase
}

type StartStreaming struct {
	RequestBase
	Stream stream.Stream `json:"stream"`
}

type StopStreaming struct {
	RequestBase
}

type SetStreamSettings struct {
	RequestBase
	Settings stream.Settings `json:"settings"`
	Save     bool            `json:"save"`
}

type SaveStreamSettings struct {
	RequestBase
}