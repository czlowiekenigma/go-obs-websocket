package events

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type rawEvent struct {
	eventType string
	streamTC  time.Duration
	recTC     time.Duration
}

type Event interface {
	UpdateType() string
	StreamTimecode() (time.Duration, bool)
	RecordTimecode() (time.Duration, bool)
	copyFromOther(e Event)
}

func (e *rawEvent) UpdateType() string {
	return e.eventType
}

func (e *rawEvent) StreamTimecode() (time.Duration, bool) {
	if e.streamTC < 0 {
		return 0, false
	}
	return e.streamTC, true
}

func (e *rawEvent) RecordTimecode() (time.Duration, bool) {
	if e.recTC < 0 {
		return 0, false
	}
	return e.recTC, true
}

func (e *rawEvent) copyFromOther(other Event) {
	e.eventType = other.UpdateType()
	var ok bool
	if e.streamTC, ok = other.StreamTimecode(); ok == false {
		e.streamTC = -1
	}

	if e.recTC, ok = other.RecordTimecode(); ok == false {
		e.recTC = -1
	}
}

func parseTC(tc string) (time.Duration, error) {
	if len(tc) == 0 {
		return -1, fmt.Errorf("obsws: invalid Timecode format '': empty string")
	}
	var h, m, s, ms int
	_, err := fmt.Sscanf(tc, "%d:%d:%d.%d", &h, &m, &s, &ms)
	if err != nil {
		return -1, fmt.Errorf("obsws: invalid Timecode format '%s': %s", tc, err)
	}

	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(ms)*time.Millisecond, nil
}

// ErrNotEventMessage is an error returned by UnmarshalEvent when the
// JSON data does not correspond to an Event message.
type ErrNotEventMessage struct{}

func (e ErrNotEventMessage) Error() string {
	return "obsws: message is not an event"
}

type ErrUnknownEventType struct {
	Type string
}

func (e ErrUnknownEventType) Error() string {
	return "obsws: unknown event type '" + e.Type + "'"
}

// UnmarshalEvent unmarshal an Event formatted in JSON. Cannot use the
// standard JSON interface as it would be too complicated to implement
func UnmarshalEvent(data []byte) (Event, error) {
	//first we extract the generic part
	aux := struct {
		StreamTCStr string `json:"stream-timecode"`
		RecTCStr    string `json:"rec-timecode"`
		UpdateType  string `json:"update-type"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, err
	}

	if len(aux.UpdateType) == 0 {
		return nil, ErrNotEventMessage{}
	}
	rawE := &rawEvent{
		eventType: aux.UpdateType,
		recTC:     -1,
		streamTC:  -1,
	}

	if len(aux.StreamTCStr) > 0 {
		var err error
		rawE.streamTC, err = parseTC(aux.StreamTCStr)
		if err != nil {
			return nil, err
		}
	}

	if len(aux.RecTCStr) > 0 {
		var err error
		rawE.recTC, err = parseTC(aux.RecTCStr)
		if err != nil {
			return nil, err
		}
	}

	// now we extract the generic part
	evType, ok := eventFactory[rawE.UpdateType()]
	if !ok {
		return nil, ErrUnknownEventType{rawE.UpdateType()}
	}

	evInst := reflect.New(evType)
	ev := evInst.Interface().(Event)

	err := json.Unmarshal(data, &ev)
	if err != nil {
		return nil, err
	}
	// copy back initial data
	ev.copyFromOther(rawE)
	return ev, nil
}

var eventFactory map[string]reflect.Type

func init() {
	eventFactory = map[string]reflect.Type{
		//Scenes
		"SwitchScenes": reflect.TypeOf(SwitchScenes{}),
		"ScenesChanged": reflect.TypeOf(ScenesChanged{}),
		"SceneCollectionChanged": reflect.TypeOf(SceneCollectionChanged{}),
		"SceneCollectionListChanged": reflect.TypeOf(SceneCollectionListChanged{}),

		//Transitions
		"SwitchTransition": reflect.TypeOf(SwitchTransition{}),
		"TransitionListChanged": reflect.TypeOf(TransitionListChanged{}),
		"TransitionDurationChanged": reflect.TypeOf(TransitionDurationChanged{}),
		"TransitionBegin": reflect.TypeOf(TransitionBegin{}),

		//Profiles
		"ProfileChanged": reflect.TypeOf(ProfileChanged{}),
		"ProfileListChanged": reflect.TypeOf(ProfileListChanged{}),

		//Streaming
		"StreamStarting": reflect.TypeOf(StreamStarting{}),
		"StreamStarted": reflect.TypeOf(StreamStarted{}),
		"StreamStopping": reflect.TypeOf(StreamStopping{}),
		"StreamStopped": reflect.TypeOf(StreamStopped{}),
		"StreamStatus": reflect.TypeOf(StreamStatus{}),

		//Recording
		"RecordingStarting": reflect.TypeOf(RecordingStarting{}),
		"RecordingStarted": reflect.TypeOf(RecordingStarted{}),
		"RecordingStopping": reflect.TypeOf(RecordingStopping{}),
		"RecordingStopped": reflect.TypeOf(RecordingStopped{}),

		//Replay buffer
		"ReplayStarting": reflect.TypeOf(ReplayStarting{}),
		"ReplayStarted": reflect.TypeOf(ReplayStarted{}),
		"ReplayStopping": reflect.TypeOf(ReplayStopping{}),
		"ReplayStopped": reflect.TypeOf(ReplayStopped{}),

		//Other
		"Exiting": reflect.TypeOf(Exiting{}),

		//General
		"Heartbeat": reflect.TypeOf(Heartbeat{}),

		//Sources
		"SourceOrderChanged": reflect.TypeOf(SourceOrderChanged{}),
		"SceneItemAdded": reflect.TypeOf(SceneItemAdded{}),
		"SceneItemRemoved": reflect.TypeOf(SceneItemRemoved{}),
		"SceneItemVisibilityChanged": reflect.TypeOf(SceneItemVisibilityChanged{}),

		//Studio mode
		"PreviewSceneChanged": reflect.TypeOf(PreviewSceneChanged{}),
		"StudioModeSwitched": reflect.TypeOf(StudioModeSwitched{}),
	}
}

