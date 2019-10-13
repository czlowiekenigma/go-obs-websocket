package client

import (
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) StartStopRecording() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StartStopRecording{}))
	return
}

func (c *Client) StartRecording() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StartRecording{}))
	return
}

func (c *Client) StopRecording() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StopRecording{}))
	return
}

func (c *Client) PauseRecording() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.PauseRecording{}))
	return
}

func (c *Client) ResumeRecording() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.ResumeRecording{}))
	return
}

func (c *Client) SetRecordingFolder(req *requests.SetRecordingFolder) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetRecordingFolder() (resp *responses.GetRecordingFolder, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetRecordingFolder{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetRecordingFolder)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}