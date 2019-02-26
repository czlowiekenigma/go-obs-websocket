package responses

import (
	"fmt"
	"strings"
)

type Response interface {
	GetError() error
	GetMessageID() string
}

type ResponseBase struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (r ResponseBase) GetError() error {
	if strings.ToLower(r.Status) == "ok" {
		return nil
	}

	return fmt.Errorf("obsws: status: %s error: %s", r.Status, r.Error)
}

func (r ResponseBase) GetMessageID() string {
	return r.MessageID
}