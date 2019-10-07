package client

import (
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) GetStreamingStatus() (resp *responses.GetStreamingStatus, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetStreamingStatus{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetStreamingStatus)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) StartStopStreaming() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StartStopStreaming{}))
	return
}

func (c *Client) StartStreaming(req *requests.StartStreaming) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) StopStreaming() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StopStreaming{}))
	return
}

func (c *Client) SetStreamingSettings(req *requests.SetStreamSettings) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetStreamSettings() (resp *responses.GetStreamSettings, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetStreamSettings{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetStreamSettings)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SaveStreamSettings() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.SaveStreamSettings{}))
	return
}
