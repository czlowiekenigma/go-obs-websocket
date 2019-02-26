package client

import (
	"fmt"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
)

func (c *Client) SetCurrentScene(req *requests.SetCurrentScene) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetCurrentScene() (resp *responses.GetCurrentScene, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetCurrentScene{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetCurrentScene)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
	return
}

func (c *Client) GetSceneList() (resp *responses.GetSceneList, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSceneList{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetSceneList)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) ReorderSceneItems(req *requests.ReorderSceneItems) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}