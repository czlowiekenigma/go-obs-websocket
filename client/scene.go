package client

import (
	"fmt"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
)

func (c *Client) SetCurrentSceneCollection(req *requests.SetCurrentSceneCollection) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetCurrentSceneCollection() (resp *responses.GetCurrentSceneCollection, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetCurrentSceneCollection{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetCurrentSceneCollection)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) ListSceneCollections() (resp *responses.ListSceneCollections, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.ListSceneCollections{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.ListSceneCollections)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}