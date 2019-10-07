package stats

type Stats struct {
	FPS                 float64 `json:"fps"`
	RenderTotalFrames   int     `json:"render-total-frames"`
	RenderMissedFrames  int     `json:"render-missed-frames"`
	OutputTotalFrames   int     `json:"output-total-frames"`
	OutputSkippedFrames int     `json:"output-skipped-frames"`
	AverageFrameTime    float64 `json:"average-frame-time"`
	CPUUsage            float64 `json:"cpu-usage"`
	MemoryUsage         float64 `json:"memory-usage"`
	FreeDiskSpace       float64 `json:"free-disk-space"`
}
