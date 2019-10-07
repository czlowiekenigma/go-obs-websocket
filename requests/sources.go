package requests

import (
	"github.com/czlowiekenigma/go-obs-websocket/scene/source"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/filter"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text/align"
	"github.com/czlowiekenigma/go-obs-websocket/scene/source/text/vertical-align"
)

type GetVolume struct {
	RequestBase
	Source string `json:"source"`
}

type SetVolume struct {
	RequestBase
	Source string  `json:"source"`
	Volume float64 `json:"volume"`
}

type GetMute struct {
	RequestBase
	Source string `json:"source"`
}

type SetMute struct {
	RequestBase
	Source string `json:"source"`
	Mute   bool   `json:"mute"`
}

type ToggleMute struct {
	RequestBase
	Source string `json:"source"`
}

type SetSyncOffset struct {
	RequestBase
	Source string `json:"source"`
	Offset int    `json:"offset"`
}

type GetSyncOffset struct {
	RequestBase
	Source string `json:"source"`
}

type GetSourceSettings struct {
	RequestBase
	Name string            `json:"sourceName"`
	Type source.SourceType `json:"sourceType,omitempty"`
}

type SetSourceSettings struct {
	RequestBase
	Name     string            `json:"sourceName"`
	Type     source.SourceType `json:"sourceType,omitempty"`
	Settings interface{}       `json:"sourceSettings"`
}

type GetTextGDIPlusProperties struct {
	RequestBase
	Source string `json:"source"`
}

type SetTextGDIPlusProperties struct {
	RequestBase
	Source        string                       `json:"source"`
	Color         int                          `json:"color"`
	Align         align.Align                  `json:"align,omitempty"`
	VerticalAlign vertical_align.VerticalAlign `json:"valign,omitempty"`
	Font          text.Font                    `json:"font,omitempty"`
	Text          string                       `json:"text,omitempty"`
	Vertical      bool                         `json:"vertical,omitempty"`
	File          string                       `json:"file,omitempty"`
	ReadFromFile  bool                         `json:"read_from_file,omitempty"`
	Render        bool                         `json:"render,omitempty"`
	text.Background
	text.ChatLog
	text.Extends
	text.Gradient
	text.Outline
}

type GetTextFreetype2Properties struct {
	RequestBase
	Source string `json:"source"`
}

type SetTextFreetype2Properties struct {
	RequestBase
	Source      string    `json:"source"`
	Color1      int       `json:"color1,omitempty"`
	Color2      int       `json:"color2,omitempty"`
	CustomWidth int       `json:"custom_width,omitempty"`
	DropShadow  bool      `json:"drop_shadow,omitempty"`
	Font        text.Font `json:"font,omitempty"`
	LogMode     bool      `json:"log_mode,omitempty"`
	Outline     bool      `json:"outline,omitempty"`
	Text        string    `json:"text,omitempty"`
	TextFile    string    `json:"text_file,omitempty"`
	WordWrap    bool      `json:"word_wrap,omitempty"`
}

type GetBrowserSourceProperties struct {
	RequestBase
	Source string `json:"source"`
}

type SetBrowserSourceProperties struct {
	RequestBase
	Source      string `json:"source"`
	IsLocalFile bool   `json:"is_local_file,omitempty"`
	LocalFile   string `json:"local_file,omitempty"`
	URL         string `json:"url,omitempty"`
	CSS         string `json:"css,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	FPS         int    `json:"fps,omitempty"`
	Shutdown    bool   `json:"shutdown,omitempty"`
}

type GetSourceFilters struct {
	RequestBase
	SourceName string `json:"sourceName"`
}

type AddFilterToSource struct {
	RequestBase
	SourceName     string `json:"sourceName"`
	FilterName     string `json:"filterName"`
	FilterType     string `json:"filterType"`
	FilterSettings string `json:"filterSettings"`
}

type RemoveFilterFromSource struct {
	RequestBase
	SourceName string `json:"sourceName"`
	FilterName string `json:"filterName"`
}

type ReorderSourceFilter struct {
	RequestBase
	SourceName string `json:"sourceName"`
	FilterName string `json:"filterName"`
	NewIndex   int    `json:"newIndex"`
}

type MoveSourceFilter struct {
	RequestBase
	SourceName   string              `json:"sourceName"`
	FilterName   string              `json:"filterName"`
	MovementType filter.MovementType `json:"movement_type"`
}

type SetSourceFilterSettings struct {
	RequestBase
	SourceName     string      `json:"sourceName"`
	FilterName     string      `json:"filterName"`
	FilterSettings interface{} `json:"filterSettings"`
}
