package client

import (
	"fmt"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
)

func (c *Client) GetSceneItemProperties(req *requests.GetSceneItemProperties) (resp *responses.GetSceneItemProperties, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSceneItemProperties{}, req))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetSceneItemProperties)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetSceneItemProperties(req *requests.SetSceneItemProperties) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) ResetSceneItem(req *requests.ResetSceneItem) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) DeleteSceneItem(req *requests.DeleteSceneItem) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) DuplicateSceneItem(req *requests.DuplicateSceneItem) (resp *responses.DuplicateSceneItem, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.DuplicateSceneItem{}, req))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.DuplicateSceneItem)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}