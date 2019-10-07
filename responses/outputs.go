package responses

import "github.com/czlowiekenigma/go-obs-websocket/output"

type ListOutputs struct {
	*ResponseBase
	Outputs []*output.Output `json:"outputs"`
}

type GetOutputInfo struct {
	*ResponseBase
	OutputInfo *output.Output `json:"outputInfo"`
}
