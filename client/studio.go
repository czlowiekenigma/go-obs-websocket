package client

import (
	"fmt"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
)

func (c *Client) GetStudioModeStatus() (resp *responses.GetStudioModeStatus, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetStudioModeStatus{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetStudioModeStatus)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetPreviewScene() (resp *responses.GetPreviewScene, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetPreviewScene{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetPreviewScene)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetPreviewScene(req *requests.SetPreviewScene) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) TransitionToProgram(req *requests.TransitionToProgram) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) EnableStudioMode() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.EnableStudioMode{}))
	return
}

func (c *Client) DisableStudioMode() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.DisableStudioMode{}))
	return
}

func (c *Client) ToggleStudioMode() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.ToggleStudioMode{}))
	return
}