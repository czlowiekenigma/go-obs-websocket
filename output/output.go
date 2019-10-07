package output

type Output struct {
	Name          string
	Type          string
	Width         int
	Flags         Flags
	Settings      interface{}
	Active        bool
	Reconnecting  bool
	Congestion    bool
	TotalFrames   int `json:"totalFrames"`
	DroppedFrames int `json:"droppedFrames"`
	TotalBytes    int `json:"totalBytes"`
}
