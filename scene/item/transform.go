package item

import "github.com/czlowiekenigma/go-obs-websocket/transforms"

type SceneItemTransform struct {
	transforms.Position
	transforms.Rotation
	transforms.Scale
	transforms.Crop
	transforms.Bounds
	Visible         bool                  `json:"visible"`
	Locked          bool                  `json:"locked"`
	SourceWidth     int                   `json:"sourceWidth"`
	SourceHeight    int                   `json:"sourceHeight"`
	Width           float64               `json:"width"`
	Height          float64               `json:"height"`
	ParentGroupName string                `json:"parentGroupName"`
	GroupChildren   []*SceneItemTransform `json:"groupChildren"`
}
