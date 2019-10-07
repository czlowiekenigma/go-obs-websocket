package scene

import "github.com/czlowiekenigma/go-obs-websocket/scene/source"

type Scene struct {
	Name    string          `json:"name"`
	Sources []source.Source `json:"sources"`
}
