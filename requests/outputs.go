package requests

type GetOutputInfo struct {
	RequestBase
	Output string `json:"outputName"`
}

type StartOutput struct {
	RequestBase
	Name string `json:"outputName"`
}

type StopOutput struct {
	RequestBase
	Name string `json:"outputName"`
}