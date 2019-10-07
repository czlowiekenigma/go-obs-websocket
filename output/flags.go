package output

type Flags struct {
	RawValue   int  `json:"rawValue"`
	Audio      bool `json:"audio"`
	Video      bool `json:"video"`
	Encoded    bool `json:"encoded"`
	MultiTrack bool `json:"multiTrack"`
	Service    bool `json:"service"`
}
