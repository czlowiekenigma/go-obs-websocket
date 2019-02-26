package requests

type Authenticate struct {
	RequestBase
	Auth string `json:"auth"`
}

type SetHeartbeat struct {
	RequestBase
	Enable bool `json:"enable"`
}

type SetFilenameFormatting struct {
	RequestBase
	FilenameFormatting string `json:"filename-formatting"`
}


