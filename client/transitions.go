package client

import (
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) GetTransitionList() (resp *responses.GetTransitionList, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetTransitionList{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetTransitionList)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetCurrentTransition() (resp *responses.GetCurrentTransition, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetTransitionList{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetCurrentTransition)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetCurrentTransition(req *requests.SetCurrentTransition) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) SetTransitionDuration(req *requests.SetTransitionDuration) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetTransitionDuration() (resp *responses.GetTransitionDuration, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetTransitionDuration{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetTransitionDuration)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}