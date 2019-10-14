package responses

import (
	"github.com/czlowiekenigma/go-obs-websocket/scene/source"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/filter"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text/align"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text/vertical-align"
)

type GetSourcesList struct {
	*ResponseBase
	Sources []source.Source `json:"sources"`
}

type GetSourceTypesList struct {
	*ResponseBase
	SourceTypes []source.SourceTypeObject `json:"types"`
}

type GetVolume struct {
	*ResponseBase
	Name   string  `json:"name"`
	Volume float64 `json:"volume"`
	Muted  bool    `json:"muted"`
}

type GetMute struct {
	*ResponseBase
	Name  string `json:"name"`
	Muted bool   `json:"muted"`
}

type GetSyncOffset struct {
	*ResponseBase
	Name   string `json:"name"`
	Offset int    `json:"offset"`
}

type GetSourceSettings struct {
	*ResponseBase
	Name     string            `json:"sourceName"`
	Type     source.SourceType `json:"sourceType"`
	Settings interface{}       `json:"sourceSettings"`
}

type SetSourceSettings struct {
	*ResponseBase
	Name string            `json:"sourceName"`
	Type source.SourceType `json:"sourceType"`
}

type GetTextGDIPlusProperties struct {
	*ResponseBase
	Source        string                       `json:"source"`
	Color         int                          `json:"color"`
	Align         align.Align                  `json:"align"`
	VerticalAlign vertical_align.VerticalAlign `json:"valign"`
	Font          text.Font                    `json:"font"`
	Text          string                       `json:"text"`
	Vertical      bool                         `json:"vertical"`
	File          string                       `json:"file"`
	ReadFromFile  bool                         `json:"read_from_file"`
	text.Background
	text.ChatLog
	text.Extends
	text.Gradient
	text.Outline
}

type GetTextFreetype2Properties struct {
	*ResponseBase
	Source      string    `json:"source"`
	Color1      int       `json:"color1"`
	Color2      int       `json:"color2"`
	CustomWidth int       `json:"custom_width"`
	DropShadow  bool      `json:"drop_shadow"`
	Font        text.Font `json:"font"`
	LogMode     bool      `json:"log_mode"`
	Outline     bool      `json:"outline"`
	Text        string    `json:"text"`
	TextFile    string    `json:"text_file"`
	WordWrap    bool      `json:"word_wrap"`
}

type GetBrowserSourceProperties struct {
	*ResponseBase
	Source      string `json:"source"`
	IsLocalFile bool   `json:"is_local_file"`
	LocalFile   string `json:"local_file"`
	URL         string `json:"url"`
	CSS         string `json:"css"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FPS         int    `json:"fps"`
	Shutdown    bool   `json:"shutdown"`
}

type GetSpecialSources struct {
	*ResponseBase
	Desktop1 string `json:"desktop1,omitempty"`
	Desktop2 string `json:"desktop2,omitempty"`
	Mic1     string `json:"mic-1"`
	Mic2     string `json:"mic-2"`
	Mic3     string `json:"mic-3"`
}

type GetSourceFilters struct {
	*ResponseBase
	Filters []*filter.Filter `json:"filters"`
}

type TakeSourceScreenshot struct {
	*ResponseBase
	SourceName string `json:"sourceName"`
	Img        string `json:"img"`
	ImageFile  string `json:"imageFile"`
}
