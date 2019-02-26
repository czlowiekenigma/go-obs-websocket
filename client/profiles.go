package client

import (
	"fmt"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
)

func (c *Client) SetCurrentProfile(req *requests.SetCurrentProfile) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetCurrentProfile() (resp *responses.GetCurrentProfile, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetCurrentProfile{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetCurrentProfile)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) ListProfiles() (resp *responses.ListProfiles, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.ListProfiles{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.ListProfiles)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}