package text

type Gradient struct {
	Gradient          bool    `json:"gradient,omitempty"`
	GradientColor     int     `json:"gradient_color,omitempty"`
	GradientDirection float64 `json:"gradient_dir,omitempty"`
	GradientOpacity   int     `json:"gradient_opacity,omitempty"`
}
