package obsws

import (
	"encoding/json"
	"fmt"
	"time"
)

// Event can be returned from the client at any time
type Event interface {
	// UpdateType specifies which kind of event it is
	UpdateType() string
	// StreamTimecode specifies the time since the stream started. It
	// return false if the stream did not yet started.
	StreamTimecode() (time.Duration, bool)
	// RecordTimecode specifies the time since the recording
	// started. It return false of the stream did not yet stated/
	RecordTimecode() (time.Duration, bool)
}

type event struct {
	EventType string `json:"update-type"`
	streamTC  time.Duration
	recTC     time.Duration
}

func (e event) UpdateType() string {
	return e.EventType
}

func (e event) StreamTimecode() (time.Duration, bool) {
	if e.streamTC < 0 {
		return 0, false
	}
	return e.streamTC, true
}

func (e event) RecordTimecode() (time.Duration, bool) {
	if e.recTC < 0 {
		return 0, false
	}
	return e.recTC, true
}

func parseTC(tc string) (time.Duration, error) {
	if len(tc) == 0 {
		return -1, fmt.Errorf("Invalid Timecode format '': empty string")
	}
	var h, m, s, ms int
	_, err := fmt.Sscanf(tc, "%d:%d:%d.%d", &h, &m, &s, &ms)
	if err != nil {
		return -1, fmt.Errorf("Invalid Timecode format '%s': %s", tc, err)
	}

	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second + time.Duration(ms)*time.Millisecond, nil
}

func (e *event) UnmarshalJSON(data []byte) error {
	type Alias event
	aux := &struct {
		StreamTCStr string `json:"stream-timecode"`
		RecTCStr    string `json:"rec-timecode"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	e.recTC = -1
	e.streamTC = -1

	if len(aux.StreamTCStr) > 0 {
		var err error
		e.streamTC, err = parseTC(aux.StreamTCStr)
		if err != nil {
			return err
		}
	}

	if len(aux.RecTCStr) > 0 {
		var err error
		e.recTC, err = parseTC(aux.RecTCStr)
		if err != nil {
			return err
		}
	}

	return nil
}
