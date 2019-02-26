package responses

import (
	"encoding/json"
	"strings"
)

type GetVersion struct {
	*ResponseBase
	Version             float64 `json:"version"`
	OBSWebsocketVersion string  `json:"obs-websocket-version"`
	OBSStudioVersion    string  `json:"obs-studio-version"`
	AvailableRequests   []string
}

func (v *GetVersion) MarshalJSON() ([]byte, error) {
	type Alias GetVersion

	return json.Marshal(&struct {
		*Alias
		AvailableRequest string `json:"available_requests"`
	}{
		Alias:            (*Alias)(v),
		AvailableRequest: strings.Join(v.AvailableRequests, ","),
	})
}

func (v *GetVersion) UnmarshalJSON(b [] byte) (err error) {
	type Alias GetVersion
	aux := &struct {
		*Alias
		AvailableRequest string `json:"available_requests"`
	}{
		Alias: (*Alias)(v),
	}

	if err = json.Unmarshal(b, &aux); err != nil {
		return
	}

	v.AvailableRequests = strings.Split(aux.AvailableRequest, ",")
	return
}

type GetAuthRequired struct {
	*ResponseBase
	AuthRequired bool   `json:"authRequired"`
	Challenge    string `json:"challenge,omitempty"`
	Salt         string `json:"salt,omitempty"`
}

type GetFilenameFormatting struct {
	ResponseBase
	FilenameFormatting string `json:"filename-formatting"`
}
