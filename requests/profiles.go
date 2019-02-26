package requests

type SetCurrentProfile struct {
	RequestBase
	ProfileName string `json:"profile-name"`
}

type GetCurrentProfile struct {
	RequestBase
}
