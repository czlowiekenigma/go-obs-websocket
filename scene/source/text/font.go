package text

type Font struct {
	Face  string `json:"face"`
	Flags int    `json:"flags"`
	Size  int    `json:"size"`
	Style string `json:"style"`
}
