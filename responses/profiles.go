package responses

type GetCurrentProfile struct {
	*ResponseBase
	ProfileName string `json:"profile-name"`
}

type ListProfiles struct {
	*ResponseBase
	Profiles []interface{} `json:"profiles"` //TODO Change interface{} to Profile
}
