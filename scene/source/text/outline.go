package text

type Outline struct {
	Outline        bool `json:"outline,omitempty"`
	OutlineColor   int  `json:"outline_color,omitempty"`
	OutlineSize    int  `json:"outline_size,omitempty"`
	OutlineOpacity int  `json:"outline_opacity,omitempty"`
}
