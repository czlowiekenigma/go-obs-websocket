package obsws

import (
	"encoding/json"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type EventSuite struct{}

var _ = Suite(&EventSuite{})

func (s *EventSuite) TestTimecodeParse(c *C) {
	tdata := map[string]time.Duration{
		"00:00:00.000": 0,
		"01:00:00.000": 1 * time.Hour,
		"00:01:00.000": 1 * time.Minute,
		"00:00:01.000": 1 * time.Second,
		"00:00:00.042": 42 * time.Millisecond,
	}

	for tc, expected := range tdata {
		res, err := parseTC(tc)
		if ok := c.Check(err, IsNil); ok == false {
			continue
		}
		c.Check(res, Equals, expected)
	}
}

func (s *EventSuite) TestExtractTimecode(c *C) {
	itdata := map[string]string{
		"Invalid Timecode format '.*':.*": "0:0",
	}

	for errorMatch, tc := range itdata {
		res, err := parseTC(tc)
		c.Check(res, Equals, time.Duration(-1))
		c.Check(err, ErrorMatches, errorMatch)
	}

}

func (s *EventSuite) TestEventExtraction(c *C) {
	tdata := map[string]event{
		`{"update-type":"foo"} `:                                  event{EventType: "foo", streamTC: -1, recTC: -1},
		`{"update-type":"foo","stream-timecode":"01:00:00.000"} `: event{EventType: "foo", streamTC: 1 * time.Hour, recTC: -1},
		`{"update-type":"foo","rec-timecode":"00:01:00.000"} `:    event{EventType: "foo", recTC: 1 * time.Minute, streamTC: -1},
	}

	for jsonData, expected := range tdata {
		res := event{}
		err := json.Unmarshal([]byte(jsonData), &res)
		c.Check(err, IsNil)
		c.Check(res, Equals, expected)
	}

	itdata := map[string]string{
		"json:.*":                           `{"update-type":42}`,
		"Invalid Timecode format '.*': .*":  `{"update-type":"foo","stream-timecode":"a"}`,
		"Invalid Timecode format 'a.*': .*": `{"update-type":"foo","rec-timecode":"a"}`,
	}

	for errorMatch, jsonData := range itdata {
		res := event{recTC: -1, streamTC: -1}
		err := json.Unmarshal([]byte(jsonData), &res)
		c.Check(err, ErrorMatches, errorMatch, Commentf(jsonData))
		c.Check(res.streamTC, Equals, time.Duration(-1))
		c.Check(res.recTC, Equals, time.Duration(-1))
	}
}
