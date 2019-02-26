package source

type Capabilities struct {
	IsAsync          bool `json:"isAsync"`          //True if source of this type provide frames asynchronously
	HasVideo         bool `json:"hasVideo"`         //True if sources of this type provide video
	HasAudio         bool `json:"hasAudio"`         //True if sources of this type provide audio
	CanInteract      bool `json:"canInteract"`      //True if interaction with this sources of this type is possible
	IsComposite      bool `json:"isComposite"`      //True if sources of this type composite one or more sub-sources
	DoNotDuplicate   bool `json:"doNotDuplicate"`   //True if sources of this type should not be fully duplicated
	DoNotSelfMonitor bool `json:"doNotSelfMonitor"` //True if sources of this type may cause a feedback loop if it's audio is monitored and shouldn't be
}
