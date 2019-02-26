package responses

type GetRecordingFolder struct {
	*ResponseBase
	RecFolder string `json:"rec-folder"`
}
