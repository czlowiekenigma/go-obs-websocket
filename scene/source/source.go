package source

type Source struct {
	Name   string     `json:"name"`   //Unique source name
	TypeID string     `json:"typeId"` //Non-unique internal source type ID
	Type   SourceType `json:"type"`   //Source type
}

type SourceTypeObject struct {
	Source
	DisplayName     string        `json:"displayName"`     //Display name of the source type
	DefaultSettings interface{}   `json:"defaultSettings"` //Default settings of this source type
	Capabilities    *Capabilities `json:"caps"`            //Source type capabilities
}
