package client

import (
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) ListOutputs() (resp *responses.ListOutputs, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.ListOutputs{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.ListOutputs)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetOutputInfo(req *requests.GetOutputInfo) (resp *responses.GetOutputInfo, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetOutputInfo{}, req))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetOutputInfo)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) StartOutput(req *requests.StartOutput) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) StopOutput(req *requests.StopOutput) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}